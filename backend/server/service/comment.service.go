package service

import (
	"forum/backend/config"
	"forum/backend/dto"
	"forum/backend/models"
	r "forum/backend/server/repositories"
	"time"

	"github.com/gofrs/uuid/v5"
)

type CommentService struct {
	CommentRepo r.CommentRepository
	ReactRepo   r.ReactionRepository
}

func (commentService *CommentService) init() {
	commentService.CommentRepo = r.CommRepo
	commentService.ReactRepo = r.ReactRepo
}

func (commentService *CommentService) NewComment(comment models.Comment) error {
	commentId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	comment.CommentId = commentId
	err = commentService.CommentRepo.SaveComment(comment)
	if err != nil {
		return err
	}
	return nil
}
func (commentService *CommentService) GetComment(commentId string) (models.Comment, error) {
	comment, err := commentService.CommentRepo.GetComment(commentId)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
func (commentService *CommentService) EditComment(updatedComment models.Comment) error {
	err := commentService.CommentRepo.UpdateComment(updatedComment)
	if err != nil {
		return err
	}
	return nil
}
func (commentService *CommentService) DeleteComment(commentID string) error {
	err := commentService.CommentRepo.RemoveComment(commentID)
	if err != nil {
		return err
	}
	return nil
}

func (commentService *CommentService) SaveCommReaction(comment models.Comment, react, userId string) error {
	reactId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	uId := uuid.FromStringOrNil(userId)
	reaction := models.Reaction{
		ReactId:   reactId,
		Reactions: react,
		ReacType:  "COMMENT",
		UserId:    uId,
		PostId:    comment.PostId,
		CommentId: comment.CommentId,
		CreatedAt: time.Now().Format(string(config.Get("TIME_FORMAT"))),
		UpdatedAt: time.Now().Format(string(config.Get("TIME_FORMAT"))),
	}
	err = commentService.ReactRepo.SaveReaction(reaction)
	return err
}

func (commentService *CommentService) GetUserCommReact(userId, commentId string) (models.Reaction, error) {
	react, err := commentService.ReactRepo.GetReactByUser(userId, commentId, "COMMENT")
	if err != nil {
		return models.Reaction{}, err
	}
	return react, nil
}

func (commentService *CommentService) GetCommentVotes(commentId string) (int,error) {
	c,err := commentService.ReactRepo.GetVotes("COMMENT", commentId)
	if err != nil {
		return c, err
	}
	return c,nil
}

func (commentService *CommentService) GetCommentsByPostId(postId string,t models.TokenData,page int) (comments []dto.CommentDTO,err error){
	comms,err := commentService.CommentRepo.GetPostComments(postId,page)
	if err != nil {
		return []dto.CommentDTO{}, err
	}

	for _, v := range comms {
		creationDate, _ := time.Parse(config.Get("TIME_FORMAT").ToString(),v.CreatedAt)
		react,_:= commentService.GetUserCommReact(t.UserId,v.CommentId.String())
		now,_ := time.Parse(config.Get("TIME_FORMAT").ToString(),time.Now().Format(config.Get("TIME_FORMAT").ToString()))
		age := now.Sub(creationDate)	
		votes,_ := commentService.GetCommentVotes(v.CommentId.String())
		comments = append(comments, dto.CommentDTO{Comment: v,Votes: votes,Age: age.String(),UserReact: react.Reactions})
	}
	return comments,nil
}