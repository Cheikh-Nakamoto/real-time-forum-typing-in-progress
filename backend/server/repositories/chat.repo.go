package repositories

import (
	db "forum/backend/database"
	opt "forum/backend/database/operators"
	q "forum/backend/database/query"
	"forum/backend/models"

	"github.com/gofrs/uuid/v5"
)

type ChatRepository struct {
	BaseRepo
}

func (r *ChatRepository) init() {
	r.DB = db.DB
	r.TableName = db.CHATS_TABLE
}

func (r *ChatRepository) SaveChat(chat models.Chat) error {

	err := r.DB.Insert(r.TableName, chat)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) UpdateChat(chat models.Chat) error {
	err := r.DB.Update(r.TableName, chat,q.WhereOption{"chat_id":opt.Equals(chat.ChatId)})
	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) DeleteChat(chatId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"cht_id": opt.Equals(chatId)})
	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) GetChat(from string,username string) (chat models.Chat, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"requester_id": opt.Equals(username)+opt.Or("recipient_id",opt.Equals(username)),"recipient_id":opt.Equals(from)+opt.Or("requester_id",opt.Equals(from))})
	if err != nil {
		return chat, err
	}
	err = row.Scan(&chat.ChatId,&chat.Requester,&chat.Recipient,&chat.LastMessageTime, &chat.CreatedAt)
	if err != nil {
		return chat, err
	}
	return chat, nil
}

func (r *ChatRepository) GetAllChats(t models.TokenData) (chats []models.Chat, err error) {
	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"user_id": opt.Equals(t.UserId)}, "created_at DESC",nil)
	if err != nil {
		return chats, err
	}
	for rows.Next() {
		var chat models.Chat
		err = rows.Scan(&chat.ChatId, &chat.CreatedAt)
		if err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

func (r *ChatRepository) GetChatMessages(chatId string) (messages []models.Message, err error) {
	rows, err := r.DB.GetAllFrom(db.MESSAGES_TABLE, q.WhereOption{"cht_id": opt.Equals(chatId)}, "created_at DESC",nil)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.MessId, &message.ChatId, &message.Sender, &message.Body, &message.CreatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r *ChatRepository) AddUserToChat(chatId uuid.UUID, userId uuid.UUID) error {
	err := r.DB.Insert(db.USERCHATS_TABLE, models.UsersChats{ChatId: chatId, UserId: userId})
	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) GetUserChats(username string) (chats []models.Chat, err error) {

	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"requester_id": opt.Equals(username) + opt.Or("recipient_id", opt.Equals(username))}, "last_message_time DESC",nil)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(&chat.ChatId,&chat.Requester,&chat.Recipient, &chat.LastMessageTime, &chat.CreatedAt)
		if err != nil {
			println(err)
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}
