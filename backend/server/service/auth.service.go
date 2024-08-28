package service

import (
	"encoding/json"
	"fmt"
	"forum/backend/encryption"
	"forum/backend/models"
	r "forum/backend/server/repositories"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo r.UserRepository
	SessRepo r.SessionRepository
}

func (authService *AuthService) init() {
	authService.UserRepo = r.UserRepo
	authService.SessRepo = r.SessRepo
}

func (authService AuthService) CreateNewUser(user *models.User) error {
	_, err := authService.UserRepo.GetUserByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("this username is already in use")
	}
	_, err = authService.UserRepo.GetUserByEmail(user.Email)
	if err == nil {
		return fmt.Errorf("this email is already in use")
	}
	userId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.UserId = userId
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPass)
	err = authService.UserRepo.SaveUser(*user)
	return err
}

func (authService *AuthService) CheckCredentials(cred map[string]string) (models.User, error) {
	identifiant := cred["identifier"]
	pass := cred["password"]
	var user models.User
	var err error
	if user, err = authService.UserRepo.GetUserByUsername(identifiant); err != nil {
		if user, err = authService.UserRepo.GetUserByEmail(identifiant); err != nil {
			return models.User{}, fmt.Errorf("invalid credentials")
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return models.User{}, fmt.Errorf("invalid credentials")
	}
	return user, nil
}

func (authService *AuthService) GenerateTokenString(data models.TokenData) (string, error) {
	toJson, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	encryptedData, err := encryption.Encrypt(string(toJson))
	if err != nil {
		return "", err
	}
	return encryptedData, nil
}

func (authService *AuthService) GetTokenData(str string) (models.TokenData, error) {
	var data models.TokenData
	toJson, err := encryption.Decrypt(str)
	if err != nil {
		return models.TokenData{}, err
	}
	err = json.Unmarshal([]byte(toJson), &data)
	if err != nil {
		return models.TokenData{}, err
	}

	return data, nil
}

func (authService *AuthService) VerifyToken(r *http.Request) (models.TokenData, error) {
	tk := r.Header.Get("auth-token")

	if tk == "" {
		tk,_ = url.QueryUnescape(r.URL.Query().Get("token")) 
		tk = strings.ReplaceAll(tk," ","+")
	}
var reader = strings.NewReader(tk)

	data := make(map[string]string)
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return models.TokenData{}, err
	}

	token := data["token"] 
	if token == "" {
		return models.TokenData{}, fmt.Errorf("missing token")
	}

	tokenData, err := authService.GetTokenData(token)
	if err != nil{
		return models.TokenData{}, err
	}

	sess, err := authService.SessRepo.GetSession(tokenData.SessId)
	if err != nil {
		return models.TokenData{}, err
	}

	if sess.Token != token || tokenData.RemoteAddr != sess.RemoteAddr {
		return models.TokenData{}, fmt.Errorf("invalid token")
	}

	return tokenData, nil
}

func (authService *AuthService) RemoveSession(sessId string) error {
	err := authService.SessRepo.DeleteSession(sessId)
	return err
}

func (authService *AuthService) GenCookieSession(w http.ResponseWriter, user models.User, r *http.Request) models.Session {
	newSessId, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return models.Session{}
	}
	var toKenData = models.TokenData{
		SessId:     newSessId.String(),
		UserId:     user.UserId.String(),
		Username:   user.Username,
		Role:       user.Role,
		SessionExp: time.Now().Add(7 * (time.Hour * 24)),
		RemoteAddr: r.RemoteAddr,
	}
	tokenStr, err := authService.GenerateTokenString(toKenData)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return models.Session{}
	}
	var newSess = models.Session{
		SessId:     newSessId,
		UserId:     user.UserId,
		Token:      tokenStr,
		ExpireAt:   time.Now().Add(7 * (time.Hour * 24)).String(),
		CreatedAt:  time.Now().String(),
		RemoteAddr: r.RemoteAddr,
	}

	return newSess
}

func (authService *AuthService) RemExistingUsrSession(userId string) {
	existingSess, _ := authService.SessRepo.GetSessionsByUserId(userId)
	for _, sess := range existingSess[1:] {
		authService.SessRepo.DeleteSession(sess.SessId.String())
	}
}
