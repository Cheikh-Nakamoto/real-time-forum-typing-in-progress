package models

import "github.com/gofrs/uuid/v5"

type UsersChats struct {
	UserId uuid.UUID	`json:"user_id"`
	ChatId uuid.UUID	`json:"chat_id"`
}