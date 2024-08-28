package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/backend/models"
	repo "forum/backend/server/repositories"
	"forum/backend/server/service"
	"io"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gofrs/uuid/v5"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		tokenData, err := service.AuthSrvice.VerifyToken(r)
		if err != nil {
			RenderErrorPage(http.StatusUnauthorized, w)
			return
		}
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
		pId := r.URL.Query().Get("postid")
		postId, err := uuid.FromString(pId)
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
		var comment models.Comment
		err = json.Unmarshal(content, &comment)
		comment.Body = strings.ReplaceAll(comment.Body, "\"", "&quot;")
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		comment.UserId, err = uuid.FromString(tokenData.UserId)
		comment.Username = tokenData.Username
		comment.PostId = postId
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		if strings.TrimSpace(comment.Body) == "" {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
		err = service.ComSrvice.NewComment(comment)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{"msg": "success"})
		}
	}
}

func EditCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		var updatedComment models.Comment
		err = json.Unmarshal(content, &updatedComment)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		updatedComment.UserId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = service.ComSrvice.EditComment(updatedComment)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:

		commentID := r.URL.Query().Get("commentid")
		if commentID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := service.ComSrvice.DeleteComment(commentID)
		if err != nil {
			fmt.Println(err)
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		} else {
			postId := r.URL.Query().Get("postid")
			tokenData, _ := authService.VerifyToken(r)
			comments, _ := service.ComSrvice.GetCommentsByPostId(postId, tokenData,0)
			data := Data{
				Username: tokenData.Username,
				Comments: comments,
			}
			tml, err := template.ParseFiles("./templates/comment.html")
			if err != nil {
				RenderErrorPage(http.StatusInternalServerError, w)
				return
			}
			tmpl := template.Must(tml, err)
			w.WriteHeader(http.StatusOK)
			err = tmpl.Execute(w, data)
			fmt.Println(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		postId := r.URL.Query().Get("postid")
		tokenData, _ := authService.VerifyToken(r)
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		comments, _ := service.ComSrvice.GetCommentsByPostId(postId, tokenData, page)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(comments)
	}
}

func CommReactHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		RenderErrorPage(http.StatusUnauthorized, w)
		return
	}
	react := r.URL.Query().Get("react")
	commentId := r.URL.Query().Get("commentid")
	if react == "" || commentId == "" {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}
	comment, err := service.ComSrvice.GetComment(commentId)
	if err != nil {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	switch r.Method {
	case http.MethodGet:
		var votes int
		if react == "LIKE" || react == "DISLIKE" {
			reaction, err := service.ComSrvice.GetUserCommReact(tokenData.UserId, commentId)
			if err != nil {
				if err == sql.ErrNoRows {
					service.ComSrvice.SaveCommReaction(comment, react, tokenData.UserId)
				}
			} else if reaction.Reactions == react {
				repo.ReactRepo.DeleteReaction(reaction.ReactId.String())
			} else {
				reaction.Reactions = react
				repo.ReactRepo.UpdateReaction(reaction)

			}
			votes, err = service.ComSrvice.GetCommentVotes(commentId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(votes)
		} else {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
	}
}
