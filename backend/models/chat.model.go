package models

import (
	"github.com/gofrs/uuid/v5"
)

type Chat struct {
	ChatId    uuid.UUID `json:"chat_id"`
	Requester string `json:"requester_id"`
	Recipient string `json:"recipient_id"`
	CreatedAt string    `json:"created_at"`
	LastMessageTime string `json:"last_message_time"`
}
