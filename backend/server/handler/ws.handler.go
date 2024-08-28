package handler

import (
	"forum/backend/ws"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		//TODO handle error here
		return
	}
	coon, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//TODO handle error here
	}
	ws.WSHub.AddClient(coon, tokenData.Username)
}
