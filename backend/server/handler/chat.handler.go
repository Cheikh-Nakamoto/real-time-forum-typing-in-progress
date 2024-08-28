package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/backend/models"
	"forum/backend/server/cors"
	repo "forum/backend/server/repositories"
	"forum/backend/server/service"
	"net/http"
	"strconv"
)


func GetStatus(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"msg": "unauthorized"})
		return
	}

	data, err := service.ChatSrvice.GetChatStatus(tokenData.Username)

	if err != nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"msg": "no chat"})
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}

func GetChatByUser(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"msg": "unauthorized"})
		return
	}

	username := r.URL.Query().Get("username")
	if _, err := repo.UserRepo.GetUserByUsername(username); err == nil {
		chat, err := repo.ChatRepo.GetChat(tokenData.Username, username)

		if err == sql.ErrNoRows {
			chat = models.Chat{
				Requester: tokenData.Username,
				Recipient: username,
			}

			err := service.ChatSrvice.NewChat(&chat)
			fmt.Println(err)
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(chat)
	}
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"msg": "unauthorized"})
		return
	}

	page,_ := strconv.Atoi(r.URL.Query().Get("page"))
	

	chatId := r.URL.Query().Get("chatId")
	messages, err := repo.MessRepo.GetChatMessages(chatId,page)

	for i, message := range messages {
		if message.Sender == tokenData.Username {
			messages[i].IsSender = true
		} else {
			messages[i].IsSender = false

		}
	}
	w.WriteHeader(200)

	if err != nil {
		json.NewEncoder(w).Encode([]models.Message{})
		return
	}
	json.NewEncoder(w).Encode(messages)
}
