package models

import (
	"fmt"
	"log"
)

type Session struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
}

func NewSession(sessionId, userId string) *Session {
	return &Session{
		SessionID: sessionId,
		UserID:    userId,
	}
}

func (s *Session) TableName() string {
	return GetSessionTableName("session")
}

func (s *Session) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (sessionid, userid) VALUES (?, ?)", s.TableName())
	_, err := DbConnection.Exec(cmd, s.SessionID, s.UserID)
	if err != nil {
		return err
	}
	return err
}

func (s *Session) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET userid = ? WHERE sessionid = ?", s.TableName())
	_, err := DbConnection.Exec(cmd, s.UserID, s.SessionID)
	if err != nil {
		return err
	}
	return err
}

func GetSession(sessionID string) (*Session, error) {
	tableName := GetUserTableName("user")
	cmd := fmt.Sprintf(`SELECT sessionid, userid FROM %s WHERE sessionID = '%s'`,
		tableName, sessionID)
	row := DbConnection.QueryRow(cmd)
	var s Session
	err := row.Scan(&s.SessionID, &s.UserID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return NewSession(s.SessionID, s.UserID), nil
}
