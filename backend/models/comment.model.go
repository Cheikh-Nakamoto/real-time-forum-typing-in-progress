package models

import (
	"github.com/gofrs/uuid/v5"
)

type Comment struct {
	CommentId uuid.UUID `json:"comment_id"`
	UserId    uuid.UUID `json:"usr_id"`
	Username  string    `json:"username"`
	PostId    uuid.UUID `json:"pst_id"`
	Body      string    `json:"body"`
	CreatedAt string    `json:"created_at"`
}
