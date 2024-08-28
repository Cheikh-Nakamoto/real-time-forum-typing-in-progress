package service

var (
	AuthSrvice AuthService
	PostSrvice PostService
	ComSrvice  CommentService
	ChatSrvice ChatService
)

func init() {
	AuthSrvice.init()
	PostSrvice.init()
	ComSrvice.init()
	ChatSrvice.init()
}
