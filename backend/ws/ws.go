package ws

import (
	"encoding/json"
	"fmt"
	"forum/backend/config"
	"forum/backend/models"
	repo "forum/backend/server/repositories"
	"log"
	"sync"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/websocket"
)

const (
	WS_JOIN_EVENT       = "join-event"
	WS_DISCONNECT_EVENT = "disconnect-event"
	WS_MESSAGE_EVENT    = "msg-event"
	WS_READ_EVENT       = "read-event"
	WS_NEW_POST_EVENT   = "new-post-event"
	WS_NEW_USER_EVENT   = "new-user-event"
	WS_TYPING_EVENT     = "typing-event"
)

type Hub struct {
	Clients           *sync.Map
	RegisterChannel   chan *WSClient
	UnRegisterChannel chan *WSClient
	SSE               chan WSPaylaod
}

type WSClient struct {
	Username    string
	WSCoon      *websocket.Conn
	OutgoingMsg chan interface{}
}

type WSPaylaod struct {
	From string
	Type string
	Data interface{}
	To   string
}

var WSHub *Hub

func init() {
	WSHub = newHub()
	go WSHub.listen()
}

func newHub() *Hub {
	return &Hub{
		Clients:           &sync.Map{},
		RegisterChannel:   make(chan *WSClient),
		UnRegisterChannel: make(chan *WSClient),
		SSE:               make(chan WSPaylaod),
	}
}

func (wsHub *Hub) listen() {
	for {
		select {
		case client := <-wsHub.RegisterChannel:
			wsHub.Clients.Store(client.Username, client)
			log.Printf("%s is connected\n", client.Username)
		case client := <-wsHub.UnRegisterChannel:
			wsHub.Clients.Delete(client.Username)
			log.Printf("%s is disconnected\n", client.Username)
		case message := <-wsHub.SSE:
			wsHub.HandleEvent(message)
		}
	}
}

func (wsHub *Hub) AddClient(coon *websocket.Conn, username string) {
	client := &WSClient{
		Username:    username,
		WSCoon:      coon,
		OutgoingMsg: make(chan interface{}),
	}

	go client.messageReader()
	go client.messageWriter()

	wsHub.RegisterChannel <- client

	var newEvent = WSPaylaod{
		From: client.Username,
		Type: WS_JOIN_EVENT,
		Data: nil,
	}

	wsHub.HandleEvent(newEvent)

}

func (wsHub *Hub) HandleEvent(eventPayload WSPaylaod) {
	switch eventPayload.Type {
	case WS_NEW_USER_EVENT:
		wsHub.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Username != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_NEW_POST_EVENT:
		wsHub.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Username != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})

	case WS_JOIN_EVENT:
		wsHub.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Username != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		wsHub.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Username != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_MESSAGE_EVENT:
		data := eventPayload.Data.(map[string]any)
		to := data["to"].(string)
		from := eventPayload.From

		var client *WSClient
		var message models.Message

		if chat, err := repo.ChatRepo.GetChat(from, to); err != nil {
			return
		} else {
			chat.LastMessageTime = time.Now().Format(config.Get("TIME_FORMAT").ToString())
			repo.ChatRepo.UpdateChat(chat)
		}

		cid := data["chatId"].(string)
		chatId := uuid.FromStringOrNil(cid)
		message = models.Message{
			Sender:    eventPayload.From,
			Body:      data["content"].(string),
			CreatedAt: data["time"].(string),
			ChatId:    chatId,
		}
		message.IsSender = false

		var event = WSPaylaod{
			Type: WS_MESSAGE_EVENT,
			From: eventPayload.From,
			Data: message,
			To:   to,
		}
		c, ok := WSHub.Clients.Load(to)
		if ok {
			client = c.(*WSClient)
			client.OutgoingMsg <- event
		}

		sender, ok := wsHub.Clients.Load(eventPayload.From)
		if ok {
			senderClient := sender.(*WSClient)
			message.IsSender = true
			event.Data = message
			senderClient.OutgoingMsg <- event
		}

		repo.MessRepo.SaveMessage(&message)

	case WS_READ_EVENT:
		data := eventPayload.Data.(map[string]any)
		chatId := data["username"].(string)
		username := eventPayload.From

		messages, err := repo.MessRepo.GetChatUnreadMessages(chatId, username)
		if err != nil {
			return
		}

		for _, message := range messages {
			message.Read = true
			err := repo.MessRepo.UpdateMessage(message)

			if err != nil {
				return
			}
		}
	case WS_TYPING_EVENT:
		//on recupere l'event typing et on le renvoie au chat correspondant
		data := eventPayload.Data.(map[string]any)
		to := data["to"].(string)
		from := eventPayload.From
		fmt.Println(data, to, from)
		var client *WSClient
		var event = WSPaylaod{
			Type: WS_TYPING_EVENT,
			From: from,
			Data: data,
			To:   to,
		}
		c, ok := WSHub.Clients.Load(to)
		if ok {
			client = c.(*WSClient)
			client.OutgoingMsg <- event
		}
	}
}

func (client *WSClient) messageReader() {
	for {
		_, message, err := client.WSCoon.ReadMessage()
		if err != nil {
			WSHub.UnRegisterChannel <- client

			var newEvent = WSPaylaod{
				From: client.Username,
				Type: WS_DISCONNECT_EVENT,
				Data: nil,
			}
			WSHub.HandleEvent(newEvent)
			return
		}
		var payload map[string]any
		err = json.Unmarshal(message, &payload)
		if err != nil {
			return
		}

		eventType := payload["type"].(string)

		wsEvent := WSPaylaod{
			From: client.Username,
			Type: eventType,
			Data: payload,
		}

		WSHub.HandleEvent(wsEvent)
	}
}

func (client *WSClient) messageWriter() {
	for message := range client.OutgoingMsg {
		data, err := json.Marshal(message)
		if err != nil {
			return
		}
		err = client.WSCoon.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return
		}
	}
}
