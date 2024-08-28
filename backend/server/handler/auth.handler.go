package handler

import (
	"encoding/json"
	"forum/backend/config"
	"forum/backend/models"
	"forum/backend/server/cors"
	"forum/backend/server/repositories"
	"forum/backend/server/service"
	"forum/backend/utils"
	"forum/backend/ws"
	"io"
	"net/http"
	"strings"
	"time"
)

var authService = service.AuthSrvice

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./frontend/index.html")
		return

	case http.MethodPost:
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
		}
		var ageGender map[string]any
		err = json.Unmarshal(content, &ageGender)
		if err != nil || ageGender["age"] == "" || ageGender["gender"] == nil {
			RenderErrorPage(http.StatusUnprocessableEntity, w)
			return
		}
		var user models.User
		err = json.Unmarshal(content, &user)
		if err != nil {
			RenderErrorPage(http.StatusUnprocessableEntity, w)
			return
		} else {
			if err := utils.VerifyUsername(user.Username); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": "wrong username"})
				return
			} else if err := utils.IsValidEmail(user.Email); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": "wrong email"})
				return
			} else if err := utils.VerifyName(user.Firstname); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": "wrong firstname"})
				return
			} else if err := utils.VerifyName(user.Lastname); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": "wrong lastname"})
				return
			}
		}
		user.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT")))
		user.UpdatedAt = user.CreatedAt

		err = authService.CreateNewUser(&user)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		} else {
			newSess := authService.GenCookieSession(w, user, r)
			newSess.UserId = user.UserId
			err := authService.SessRepo.SaveSession(newSess)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			}

			reformatedUserInfo := map[string]any{
				"username":  user.Username,
				"firstname": user.Firstname,
				"lastname":  user.Lastname,
				"email":     user.Email,
			}

			var newEvent = ws.WSPaylaod{
				From: user.Username,
				Type: ws.WS_NEW_USER_EVENT,
				Data: map[string]any{
					"username":     user.Username,
					"status":       "offline",
					"unread_count": 0,
				},
			}
			ws.WSHub.HandleEvent(newEvent)

			json.NewEncoder(w).Encode(map[string]any{"authToken": newSess.Token, "msg": "success", "user": reformatedUserInfo})

		}
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./frontend/index.html")
		return
		// http.Redirect(w,r,"/",http.StatusPermanentRedirect)
	case http.MethodPost:
		var credentials = make(map[string]string, 2)
		if content, err := io.ReadAll(r.Body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			return
		} else {

			err := json.Unmarshal(content, &credentials)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			}
			identifiant := strings.TrimSpace(credentials["identifier"])
			credentials["identifier"] = identifiant
			if err := utils.VerifyUsername(identifiant); err != nil && !strings.Contains(identifiant, "@") {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			} else if err := utils.IsValidEmail(identifiant); err != nil && strings.Contains(identifiant, "@") {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
				return
			}
		}

		user, err := authService.CheckCredentials(credentials)
		if err != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
			// fmt.Fprintln(w, err)
			return
		}
		newSess := authService.GenCookieSession(w, user, r)
		authService.SessRepo.SaveSession(newSess)

		authService.RemExistingUsrSession(user.UserId.String())
		w.WriteHeader(200)
		reformatedUserInfo := map[string]any{
			"username":  user.Username,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"email":     user.Email,
		}

		json.NewEncoder(w).Encode(map[string]any{"authToken": newSess.Token, "msg": "success", "user": reformatedUserInfo})
	default:
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	switch r.Method {
	case http.MethodDelete:

		tokenvalue, err := authService.VerifyToken(r)
		if err != nil {
			RenderErrorPage(http.StatusUnauthorized, w)
			return
		}

		authService.SessRepo.DeleteSession(tokenvalue.SessId)

		if c, ok := ws.WSHub.Clients.Load(tokenvalue.Username); ok {
			client := c.(*ws.WSClient)
			ws.WSHub.UnRegisterChannel <- client

			var newEvent = ws.WSPaylaod{
				From: client.Username,
				Type: ws.WS_DISCONNECT_EVENT,
				Data: nil,
			}
			ws.WSHub.HandleEvent(newEvent)
		}

		json.NewEncoder(w).Encode(map[string]any{"msg": "success"})

	default:
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
}

func VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	email := r.URL.Query().Get("email")
	_, err := authService.UserRepo.GetUserByEmail(email)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"msg": "email already in use"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "valid"})

}

func VerifyUsernameHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	username := r.URL.Query().Get("username")
	_, err := authService.UserRepo.GetUserByUsername(username)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"msg": "username already in use"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "valid"})
}

func VerifySessionHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]any{"msg": "unauthorized"})
		return
	}

	user, _ := repositories.UserRepo.GetUserByUsername(tokenData.Username)

	w.WriteHeader(200)
	reformatedUserInfo := map[string]any{
		"username":  user.Username,
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
		"email":     user.Email,
	}

	json.NewEncoder(w).Encode(map[string]any{"msg": "success", "user": reformatedUserInfo})

}
