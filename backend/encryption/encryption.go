package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"forum/backend/config"
)

func Encrypt(text string) (string, error) {
	secretKey := config.Get("TOKEN_SECRET").ToString()

	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(text), nil)
	textB64 := base64.StdEncoding.EncodeToString(cipherText)
	return textB64, nil
}

func Decrypt(encrtedText string) (string, error) {
	text, err := base64.StdEncoding.DecodeString(encrtedText)
	if err != nil {
		return "", err
	}
	secretKey := config.Get("TOKEN_SECRET").ToString()

	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encrtedText) <= nonceSize {
		return "", fmt.Errorf("not a valid token")
	}
	nonce, cipherText := text[:nonceSize], text[nonceSize:]
	decryptedText, err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		return "", err
	}
	return string(decryptedText), nil
}
