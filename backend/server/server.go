package server

import (
	"fmt"
	"forum/backend/config"
	db "forum/backend/database"
	"forum/backend/server/routes"
	"net/http"
)

type Server struct {
	Host string
	Port int
	DB   *db.Database
}

var (
	Servr Server
)

func New() *Server {
	return &Servr
}

func init() {
	Servr.Host = config.Get("SERVER_HOST").ToString()
	Servr.Port = config.Get("SERVER_PORT").ToInt()
	Servr.DB = db.DB
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%v", s.Port)
	fmt.Printf("server running on port %v\nhttp://%v:%v\n", s.Port, s.Host, s.Port)
	routes := routes.Route()
	http.ListenAndServe(addr, routes)
}
