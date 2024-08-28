package repositories

import (
	"database/sql"
	"fmt"
	db "forum/backend/database"
	opt "forum/backend/database/operators"
	q "forum/backend/database/query"
	"forum/backend/models"
)

type SessionRepository struct {
	BaseRepo
}

func (s *SessionRepository) init() {
	s.DB = db.DB
	s.TableName = db.SESSIONS_TABLE
}

func (s *SessionRepository) SaveSession(sess models.Session) error {
	err := s.DB.Insert(s.TableName, sess)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionRepository) UpdateSession(sess models.Session) error {
	err := s.DB.Update(s.TableName, sess, q.WhereOption{"sess_id": opt.Equals(sess.SessId)})
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionRepository) DeleteSession(sessId string) error {
	err := s.DB.Delete(s.TableName, q.WhereOption{"sess_id": opt.Equals(sessId)})
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionRepository) GetSession(sessId string) (models.Session, error) {
	var sess models.Session
	row, err := s.DB.GetOneFrom(s.TableName, q.WhereOption{"sess_id": opt.Equals(sessId)})
	if err != nil {
		return sess, err
	}
	err = row.Scan(&sess.SessId, &sess.UserId, &sess.ExpireAt, &sess.Token, &sess.CreatedAt, &sess.RemoteAddr)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Session{}, fmt.Errorf("no session found with this id")
		} else {
			fmt.Println(err)
			return models.Session{}, err
		}

	}
	return sess, nil
}

func (s *SessionRepository) GetSessionsByUserId(userId string) (sessions []models.Session, err error) {
	var sess models.Session
	rows, err := s.DB.GetAllFrom(s.TableName, q.WhereOption{"user_id": opt.Equals(userId)}, "created_at DESC",nil)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&sess.SessId, &sess.UserId, &sess.ExpireAt, &sess.Token, &sess.CreatedAt, &sess.RemoteAddr)
		sessions = append(sessions, sess)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no sessions found with this userId")
		} else {
			fmt.Println(err)
			return nil, err
		}

	}
	return sessions, nil
}
