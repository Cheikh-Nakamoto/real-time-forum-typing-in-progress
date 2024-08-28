package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/backend/config"
	"forum/backend/dto"
	"forum/backend/models"
	"forum/backend/server/cors"
	"forum/backend/server/repositories"
	"forum/backend/server/service"
	"forum/backend/utils"
	"forum/backend/ws"
	"io"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/gofrs/uuid/v5"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RenderErrorPage(http.StatusUnauthorized, w)
		return
	}
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}

		body := r.Body
		if err != nil {
			if err == io.EOF {
				RenderErrorPage(http.StatusBadRequest, w)
				return
			} else {
				fmt.Println(err)
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
		}
		var post models.Post
		post.Username = tokenData.Username

		err = json.NewDecoder(body).Decode(&post)

		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		post.Body = strings.ReplaceAll(post.Body, "\"", `&quot;`)
		post.Title = strings.ReplaceAll(post.Title, "\"", `&quot;`)

		if strings.TrimSpace(post.Body) == "" || strings.TrimSpace(post.Title) == "" {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}

		post.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		err = service.PostSrvice.NewPost(&post)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		} else {
			postDto := dto.PostDTO{
				Post:          post,
				Votes:         0,
				CommentsCount: 0,
			}
			newEvent := ws.WSPaylaod{
				Type: ws.WS_NEW_POST_EVENT,
				Data: postDto,
				From: tokenData.Username,
			}
			ws.WSHub.SSE <- newEvent

			json.NewEncoder(w).Encode(map[string]any{"msg": "success", "post": postDto})
			return
		}
	}
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, _ := service.AuthSrvice.VerifyToken(r)

	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:

		body := r.Body
		if _, err := io.ReadAll(body); err != nil {
			if err == io.EOF {
				RenderErrorPage(http.StatusBadRequest, w)
				return
			} else {
				fmt.Println(err)
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
		}
		var post models.Post

		err := json.NewDecoder(body).Decode(&post)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		post.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		err = service.PostSrvice.NewPost(&post)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		http.Redirect(w, r, "/HELLO", http.StatusPermanentRedirect)
		return
	}
}

func PostReactHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		RenderErrorPage(http.StatusUnauthorized, w)
		return
	}
	react := r.URL.Query().Get("react")
	postId := r.URL.Query().Get("postid")
	if react == "" || postId == "" {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}
	post, err := service.PostSrvice.GetPost(postId)
	if err != nil {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	switch r.Method {
	case http.MethodGet:
		var votes int
		if react == "LIKE" || react == "DISLIKE" {
			reaction, err := service.PostSrvice.GetUserPostReact(tokenData.UserId, postId)
			if err != nil {
				if err == sql.ErrNoRows {
					service.PostSrvice.SavePostReaction(post, react, tokenData.UserId)
				}
			} else if reaction.Reactions == react {
				repositories.ReactRepo.DeleteReaction(reaction.ReactId.String())
			} else {
				reaction.Reactions = react
				repositories.ReactRepo.UpdateReaction(reaction)

			}
			votes, err = service.PostSrvice.GetPostVotes(postId)
			if err != nil {
				RenderErrorPage(http.StatusInternalServerError, w)
			}
			json.NewEncoder(w).Encode(map[string]any{"votes": votes, "msg": "success"})
		} else {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
	}
}

func GetAllPostHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	switch r.Method {
	case http.MethodGet:
		tokenData, err := authService.VerifyToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"msg": "unauthorized"})
			return
		}
		pageNum := r.URL.Query().Get("page")

		liked := r.URL.Query().Get("liked")
		order := r.URL.Query().Get("order")
		category := r.URL.Query().Get("category")
		created := r.URL.Query().Get("created")
		commented := r.URL.Query().Get("commented")

		var options = map[string]string{
			"liked":     liked,
			"order":     order,
			"category":  category,
			"created":   created,
			"commented": commented,
			"page":      pageNum,
		}

		var posts []dto.PostDTO

		posts, _ = service.PostSrvice.GetAllPosts(tokenData, options)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(posts)

	}
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	tokenData, _ := authService.VerifyToken(r)
	postId := r.URL.Query().Get("postid")
	if id := uuid.FromStringOrNil(postId); id == uuid.Nil {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}
	post, err := service.PostSrvice.GetPost(postId)
	if err != nil {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	creationDate, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), post.CreatedAt)
	now, _ := time.Parse(config.Get("TIME_FORMAT").ToString(), time.Now().Format(config.Get("TIME_FORMAT").ToString()))
	age := utils.FormatDuration(now.Sub(creationDate))
	commCout, _ := service.ComSrvice.CommentRepo.GetCommentsCount(postId)
	react, _ := service.PostSrvice.GetUserPostReact(tokenData.UserId, post.PostId.String())
	votes, _ := service.PostSrvice.GetPostVotes(postId)
	categories, _ := service.PostSrvice.PostRepo.GetPostCategories(post.PostId.String())

	data.Posts = []dto.PostDTO{{Post: post, CommentsCount: commCout, Age: age, UserReact: react.Reactions, Votes: votes, Categories: categories}}
	tml, err := template.ParseFiles("templates/post.html")
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	tmpl := template.Must(tml, err)
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
