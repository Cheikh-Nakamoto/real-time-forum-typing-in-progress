package models

import "github.com/gofrs/uuid/v5"

type Message struct {
	MessId    uuid.UUID `json:"message_id"`
	ChatId    uuid.UUID `json:"cht_id"`
	Sender    string    `json:"sender_id"`
	Body      string    `json:"content"`
	CreatedAt string    `json:"created_at"`
	Read      bool      `json:"read"`
	IsSender  bool      `json:"isSender"`
}
