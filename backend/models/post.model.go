package models

import (
	"github.com/gofrs/uuid/v5"
)

type Post struct {
	PostId     uuid.UUID `json:"post_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	UserId     uuid.UUID `json:"user_id"`
	Username   string    `json:"username"`
	Status     string    `json:"status"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
}
