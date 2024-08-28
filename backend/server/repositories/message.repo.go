package repositories

import (
	db "forum/backend/database"
	opt "forum/backend/database/operators"
	q "forum/backend/database/query"
	"forum/backend/models"

	"github.com/gofrs/uuid/v5"
)

type MessageRepository struct {
	BaseRepo
}

func (r *MessageRepository) init() {
	r.DB = db.DB
	r.TableName = db.MESSAGES_TABLE
}

func (r *MessageRepository) SaveMessage(message *models.Message) error {
	msgId, _ := uuid.NewV4()
	message.MessId = msgId
	type messTable struct {
		MessId    uuid.UUID `json:"message_id"`
		ChatId    uuid.UUID `json:"cht_id"`
		Sender    string    `json:"sender_id"`
		Body      string    `json:"content"`
		CreatedAt string    `json:"created_at"`
		Read      bool      `json:"read"`
	}
	var mess = messTable{
		MessId:    message.MessId,
		Sender:    message.Sender,
		Body:      message.Body,
		CreatedAt: message.CreatedAt,
		ChatId:    message.ChatId,
		Read: false,
	}
	err := r.DB.Insert(r.TableName, mess)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteMessage(messageId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"mess_id": opt.Equals(messageId)})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) UpdateMessage(message models.Message) error {
	type messTable struct {
		MessId    uuid.UUID `json:"message_id"`
		ChatId    uuid.UUID `json:"cht_id"`
		Sender    string    `json:"sender_id"`
		Body      string    `json:"content"`
		CreatedAt string    `json:"created_at"`
		Read      bool      `json:"read"`
	}
	var mess = messTable{
		MessId:    message.MessId,
		Sender:    message.Sender,
		Body:      message.Body,
		CreatedAt: message.CreatedAt,
		ChatId:    message.ChatId,
		Read: message.Read,
	}
	err := r.DB.Update(r.TableName, mess, q.WhereOption{"message_id": opt.Equals(message.MessId)})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) GetMessage(messageId string) (message models.Message, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"mess_id": opt.Equals(messageId)})
	if err != nil {
		return message, err
	}
	err = row.Scan(&message.MessId, &message.ChatId, &message.Sender, &message.Body, &message.Read, &message.CreatedAt)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (r *MessageRepository) GetChatMessages(ChatId string,page int) (messages []models.Message, err error) {
	limit := 10
	dataOffset := page*limit

	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"cht_id": opt.Equals(ChatId)}, "created_at DESC",[]int{dataOffset,limit})
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.MessId, &message.Body, &message.ChatId, &message.Sender, &message.Read, &message.CreatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r *MessageRepository) GetChatUnreadMessages(to string, username string) (messages []models.Message, err error) {
	chat,err := ChatRepo.GetChat(username,to)
	if err!= nil {
        return messages, err
    }
	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"cht_id": opt.Equals(chat.ChatId),"sender_id":opt.NotEqual(username),"read":opt.Equals(false)}, "created_at DESC",nil)
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.MessId, &message.Body, &message.ChatId, &message.Sender, &message.Read, &message.CreatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r *MessageRepository) GetChatUnreadMessagesCount(ChatId string, username string) (count int, err error) {
	row, err := r.DB.GetCount(r.TableName, q.WhereOption{"cht_id": opt.Equals(ChatId), "sender_id": opt.NotEqual(username),"read":opt.Equals(false)})
	if err != nil {
		return 0, err
	}
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}