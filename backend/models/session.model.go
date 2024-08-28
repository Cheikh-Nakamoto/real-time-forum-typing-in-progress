package models

import (
	"github.com/gofrs/uuid/v5"
)

type Session struct {
	SessId     uuid.UUID `json:"sess_id"`
	UserId     uuid.UUID `json:"user_id"`
	ExpireAt   string    `json:"expire_at"`
	Token      string    `json:"token"`
	CreatedAt  string    `json:"created_at"`
	RemoteAddr string    `json:"remote_addr"`
}
