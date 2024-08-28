package dto

import "forum/backend/models"

type PostDTO struct {
	models.Post
	Votes         int
	Age           string
	CommentsCount int
	UserReact     string
	Categories    []models.Category
}

type CommentDTO struct {
	models.Comment
	Age       string
	Votes     int
	UserReact string
}
