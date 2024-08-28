package models

import (
	"github.com/gofrs/uuid/v5"
)

type Reaction struct {
	ReactId   uuid.UUID `json:"react_id"`
	Reactions string    `json:"reactions"`
	PostId    uuid.UUID `json:"pst_id"`
	CommentId uuid.UUID `json:"comment_id"`
	ReacType  string    `json:"react_type"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	UserId    uuid.UUID `json:"usr_id"`
}
