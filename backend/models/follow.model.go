package models

import (
	"github.com/gofrs/uuid/v5"
)

type Follow struct {
	FollowId        uuid.UUID `json:"follow_id"`
	FollowedUserId  uuid.UUID `json:"followed_user_id"`
	FollowingUserId uuid.UUID `json:"following_user_id"`
	CreatedAt       string    `json:"created_at"`
}
