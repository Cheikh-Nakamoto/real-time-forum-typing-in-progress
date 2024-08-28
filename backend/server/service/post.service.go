package service

import (
	"fmt"
	"forum/backend/config"
	"forum/backend/dto"
	"forum/backend/models"
	r "forum/backend/server/repositories"
	"forum/backend/utils"
	"time"

	"github.com/gofrs/uuid/v5"
)

type PostService struct {
	PostRepo    r.PostRepository
	ReactRepo   r.ReactionRepository
	CommentRepo r.CommentRepository
}

func (postService *PostService) init() {
	postService.PostRepo = r.PostRepo
	postService.ReactRepo = r.ReactRepo
	postService.CommentRepo = r.CommRepo
}

func (postService *PostService) NewPost(post *models.Post) error {
	postId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	post.PostId = postId
	post.Status = "PUBLISHED"
	post.CreatedAt = time.Now().Format(config.Get("TIME_FORMAT").ToString())
	post.UpdatedAt = post.CreatedAt

	err = postService.PostRepo.SavePost(*post)
	if err != nil {
		return err
	}
	return nil
}

func (postService *PostService) UpdatePost(post models.Post) error {
	err := postService.PostRepo.UpdatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (postService *PostService) RemovePost(postId string) error {
	err := postService.PostRepo.DeletePost(postId)
	if err != nil {
		return err
	}
	return nil
}

func (postService *PostService) GetPost(postId string) (models.Post, error) {
	post, err := postService.PostRepo.GetPost(postId)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (postService *PostService) FilterPostByFollow(userId string) ([]models.Post, error) {
	posts, err := postService.PostRepo.GetPostByFollow(userId)
	if err != nil {
		return []models.Post{}, err
	}
	return posts, nil
}

func (postService *PostService) SavePostReaction(post models.Post, react, userId string) error {
	reactId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	uId := uuid.FromStringOrNil(userId)
	reaction := models.Reaction{
		ReactId:   reactId,
		Reactions: react,
		ReacType:  "POST",
		UserId:    uId,
		PostId:    post.PostId,
		CreatedAt: time.Now().Format(string(config.Get("TIME_FORMAT"))),
		UpdatedAt: time.Now().Format(string(config.Get("TIME_FORMAT"))),
	}
	err = postService.ReactRepo.SaveReaction(reaction)
	return err
}

func (postService *PostService) UpdatePostReaction(reactId string, react string) error {
	reaction, err := postService.ReactRepo.GetReactById(reactId)
	if err != nil {
		return err
	}
	reaction.Reactions = react
	err = postService.ReactRepo.UpdateReaction(reaction)
	return err
}

func (postService *PostService) GetPostVotes(postId string) (int, error) {
	c, err := postService.ReactRepo.GetVotes("POST", postId)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (postService *PostService) GetUserPostReact(userId, postId string) (models.Reaction, error) {
	react, err := postService.ReactRepo.GetReactByUser(userId, postId, "POST")
	if err != nil {
		return models.Reaction{}, err
	}
	return react, nil
}

func (postService *PostService) GetAllPosts(t models.TokenData, options map[string]string) (res []dto.PostDTO, err error) {
	posts, _ := postService.PostRepo.GetPosts(t, options)
	for _, post := range posts {
		creationDate, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), post.CreatedAt)
		categories, _ := postService.PostRepo.GetPostCategories(post.PostId.String())
		react, _ := postService.GetUserPostReact(t.UserId, post.PostId.String())
		now, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), time.Now().Format(config.Get("TIME_FORMAT").ToString()))
		age := utils.FormatDuration(now.Sub(creationDate))
		votes, _ := postService.GetPostVotes(post.PostId.String())
		commentCounts, _ := postService.CommentRepo.GetCommentsCount(post.PostId.String())

		res = append(res, dto.PostDTO{Votes: votes, Post: post, CommentsCount: commentCounts, Age: age, UserReact: react.Reactions, Categories: categories})
	}
	return res, nil
}

func (postService *PostService) GetPostByKeywords(keywords []string, t models.TokenData) (res []dto.PostDTO) {
	rows := r.PostRepo.SearchSuggestions(keywords)
	var post models.Post
	var relevence_score int
	for rows.Next() {
		err := rows.Scan(&post.PostId, &post.Title, &post.Body, &post.Username, &post.UserId, &post.Status, &post.CreatedAt, &post.UpdatedAt, &relevence_score)
		if err != nil {
			fmt.Println(err)
		}
		creationDate, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), post.CreatedAt)
		categories, _ := postService.PostRepo.GetPostCategories(post.PostId.String())
		react, _ := postService.GetUserPostReact(t.UserId, post.PostId.String())
		now, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), time.Now().Format(config.Get("TIME_FORMAT").ToString()))
		age := utils.FormatDuration(now.Sub(creationDate))
		votes, _ := postService.GetPostVotes(post.PostId.String())
		commentCounts, _ := postService.CommentRepo.GetCommentsCount(post.PostId.String())

		res = append(res, dto.PostDTO{Post: post, Categories: categories, UserReact: react.Reactions, Age: age, Votes: votes, CommentsCount: commentCounts})
	}
	return res
}
