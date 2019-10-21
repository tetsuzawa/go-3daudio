package models

import (
	"fmt"
	"log"
)

type Session struct {
	SessionID string `json:"session_id"`
	UserName  string `json:"user_id"`
}

func NewSession(sessionId, userName string) *Session {
	return &Session{
		SessionID: sessionId,
		UserName:  userName,
	}
}

func (s *Session) TableName() string {
	return GetSessionTableName("session")
}

func (s *Session) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (sessionid, username) VALUES (?, ?)", s.TableName())
	_, err := DbConnection.Exec(cmd, s.SessionID, s.UserName)
	if err != nil {
		return err
	}
	return err
}

func (s *Session) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET username = ? WHERE sessionid = ?", s.TableName())
	_, err := DbConnection.Exec(cmd, s.UserName, s.SessionID)
	if err != nil {
		return err
	}
	return err
}

func GetSession(sessionID string) (*Session, error) {
	tableName := GetUserTableName("user")
	cmd := fmt.Sprintf(`SELECT sessionid, username FROM %s WHERE sessionID = '%s'`,
		tableName, sessionID)
	row := DbConnection.QueryRow(cmd)
	var s Session
	err := row.Scan(&s.SessionID, &s.UserName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return NewSession(s.SessionID, s.UserName), nil
}
