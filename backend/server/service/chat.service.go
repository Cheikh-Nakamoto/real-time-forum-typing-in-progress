package service

import (
	"forum/backend/config"
	"forum/backend/database/operators"
	"forum/backend/database/query"
	"forum/backend/models"
	"forum/backend/server/repositories"
	r "forum/backend/server/repositories"
	"forum/backend/ws"
	"slices"
	"time"

	"github.com/gofrs/uuid/v5"
)

type ChatService struct {
	ChatRepo r.ChatRepository
}

func (chatService *ChatService) init() {
	chatService.ChatRepo = r.ChatRepo
}

func (chatService *ChatService) NewChat(chat *models.Chat) error {
	chatId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	chat.ChatId = chatId
	chat.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT").ToString()))
	err = chatService.ChatRepo.SaveChat(*chat)
	if err != nil {
		return err
	}
	return nil
}

func (chatService *ChatService) UpdateChat(chat models.Chat) error {
	err := chatService.ChatRepo.UpdateChat(chat)
	if err != nil {
		return err
	}
	return nil
}

func (chatService *ChatService) DeleteChat(chatId string) error {
	err := chatService.ChatRepo.DeleteChat(chatId)
	if err != nil {
		return err
	}
	return nil
}

func (chatService *ChatService) GetAllChats(t models.TokenData) ([]models.Chat, error) {
	chats, err := chatService.ChatRepo.GetAllChats(t)
	if err != nil {
		return chats, err
	}
	return chats, nil
}

func (chatService *ChatService) GetChatMessages(chatId string) ([]models.Message, error) {
	messages, err := chatService.ChatRepo.GetChatMessages(chatId)
	if err != nil {
		return messages, err
	}
	return messages, nil
}

func (chatService *ChatService) GetChatStatus(username string) (any, error) {
	chats, err := repositories.ChatRepo.GetUserChats(username)
	if err != nil {
		println(err)
		return nil, err
	}

	type reformatedUserData struct {
		Username    string `json:"username"`
		Status      string `json:"status"`
		UnreadCount int    `json:"unread_count"`
	}

	var data []reformatedUserData

	for _, chat := range chats {
		unreadCount, _ := repositories.MessRepo.GetChatUnreadMessagesCount(chat.ChatId.String(), username)
		var uname string
		if username == chat.Recipient {
			uname = chat.Requester
		} else {
			uname = chat.Recipient
		}
		var status = "offline"
		_, ok := ws.WSHub.Clients.Load(uname)
		if ok {
			status = "online"
		}
		msgCount,_ := chatService.GetMessageCount(chat.ChatId.String())

		if msgCount == 0 {
			continue
		}
		data = append(data, reformatedUserData{Username: uname, Status: status, UnreadCount: unreadCount})
	}

	users, err := repositories.UserRepo.GetAllUsers()

	if err != nil {

		return nil, err
	}

	for _, user := range users {
		if username == user.Username {
			continue
		}
		if slices.ContainsFunc(data, func(rud reformatedUserData) bool {
			return rud.Username == user.Username
		}) {
			continue
		}
		var status = "offline"

		if _, ok := ws.WSHub.Clients.Load(user.Username); ok {
			status = "online"
		}


		data = append(data, reformatedUserData{Username: user.Username, Status: status})
	}

	return data, nil
}

func (chatService *ChatService) GetMessageCount(chatId string) (count int, err error) {
	row, err := chatService.ChatRepo.DB.GetCount("messages", query.WhereOption{"cht_id": operators.Equals(chatId)})
	if err != nil {
		return 0, err
	}
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
