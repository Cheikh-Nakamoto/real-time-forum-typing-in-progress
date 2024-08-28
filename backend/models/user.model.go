package models

import (
	"github.com/gofrs/uuid/v5"
)

type User struct {
	UserId         uuid.UUID `json:"user_id"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	Username       string    `json:"username"`
	AvatarUrl      string    `json:"avatar_url"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Status         string    `json:"status"`
	Blocked        bool      `json:"blocked"`
	EmailConfirmed bool      `json:"email_confirmed"`
	Role           string    `json:"role"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}
