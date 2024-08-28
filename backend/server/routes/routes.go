package routes

import (
	"forum/backend/server/handler"
	"net/http"
)

const (
	HOME_ENDPOINT            = "/"
	SIGNUP_ENDPOINT          = "/auth/signup"
	SIGNIN_ENDPOINT          = "/auth/signin"
	LOGOUT_ENDPOINT          = "/auth/signout"
	CREATE_POST_ENDPOINT     = "/post/create"
	STATIC_ENDPOINT          = "/static/"
	POSTREACT_ENDPOINT       = "/post/react"
	CREATE_COMMENT_ENDPOINT  = "/post/comment/create"
	GET_COMMENT_ENDPOINT     = "/post/comments"
	COMMENT_REACT_ENDPOINT   = "/comment/react"
	DELETE_COMMENT_ENDPOINT  = "/comment/delete"
	GET_POSTS_ENDPOINT       = "/posts/get"
	VERIFY_EMAIL_ENDPOINT    = "/verify/email"
	VERIFY_USERNAME_ENDPOINT = "/verify/username"
	SEARCH_SUGG_HANDLER      = "/search/sugg"
	SEARCH_ENDPOINT          = "/search"
	GET_POST_ENDPOINT        = "/post"
	ABOUT_ENDPOINT           = "/about"
	VERIF_SESS_ENDPOINT = "/verifsess"
	CHATS_ENDPOINT = "/chats"
	USER_CHAT_ENDPOINT = "/chat"
	MESSAGE_ENDPOINT = "/messages"
	USERS_STATUS_ENDPOINT = "/users-status"
	WS_ENDPOINT = "/ws"

)

func Route() *http.ServeMux {

	var mux = http.ServeMux{}
	
	mux.HandleFunc(HOME_ENDPOINT, handler.IndexHandler)
	mux.HandleFunc(SIGNUP_ENDPOINT, handler.SignUpHandler)
	mux.HandleFunc(SIGNIN_ENDPOINT, handler.SignInHandler)
	mux.HandleFunc(GET_COMMENT_ENDPOINT,handler.Authorization( handler.GetCommentsHandler))
	mux.HandleFunc(CREATE_POST_ENDPOINT, handler.Authorization(handler.CreatePostHandler))
	mux.HandleFunc(POSTREACT_ENDPOINT, handler.Authorization(handler.PostReactHandler))
	mux.HandleFunc(CREATE_COMMENT_ENDPOINT, handler.Authorization(handler.CreateCommentHandler))
	mux.HandleFunc(COMMENT_REACT_ENDPOINT, handler.Authorization(handler.CommReactHandler))
	mux.HandleFunc(LOGOUT_ENDPOINT, handler.Authorization(handler.SignOutHandler) )
	mux.HandleFunc(DELETE_COMMENT_ENDPOINT, handler.DeleteCommentHandler)
	mux.HandleFunc(GET_POSTS_ENDPOINT, handler.Authorization(handler.GetAllPostHandler) )
	mux.HandleFunc(VERIFY_EMAIL_ENDPOINT, handler.VerifyEmailHandler)
	mux.HandleFunc(VERIFY_USERNAME_ENDPOINT, handler.VerifyUsernameHandler)
	mux.HandleFunc(SEARCH_SUGG_HANDLER, handler.SearchSuggestionHandler)
	mux.HandleFunc(GET_POST_ENDPOINT, handler.GetPostHandler)
	mux.HandleFunc(SEARCH_ENDPOINT, handler.SearchHandler)
	mux.HandleFunc(ABOUT_ENDPOINT, handler.AboutHandler)
	mux.HandleFunc(VERIF_SESS_ENDPOINT, handler.VerifySessionHandler)
	mux.HandleFunc(WS_ENDPOINT, handler.WSHandler)
	mux.HandleFunc(USER_CHAT_ENDPOINT, handler.GetChatByUser)
	mux.HandleFunc(USERS_STATUS_ENDPOINT, handler.GetStatus)
	mux.HandleFunc(MESSAGE_ENDPOINT, handler.GetMessages)

	return &mux
}
