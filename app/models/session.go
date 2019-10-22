package models

import (
	"fmt"
	"log"
	"time"
)

type Session struct {
	SessionID string    `json:"session_id"`
	UserName  string    `json:"user_id"`
	Time      time.Time `json:"time"`
}

func NewSession(sessionId, userName string, timeDate time.Time) *Session {
	return &Session{
		SessionID: sessionId,
		UserName:  userName,
		Time:      timeDate,
	}
}

func (s *Session) TableName() string {
	return GetSessionTableName("session")
}

func (s *Session) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (sessionid, username, time) VALUES (?, ?, ?)", s.TableName())
	_, err := DbConnection.Exec(cmd, s.SessionID, s.UserName, s.Time.Format(time.RFC3339))
	if err != nil {
		return err
	}
	return err
}

func (s *Session) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET username = ?, time = ? WHERE sessionid = ?", s.TableName())
	_, err := DbConnection.Exec(cmd, s.UserName, s.Time.Format(time.RFC3339), s.SessionID)
	if err != nil {
		return err
	}
	return err
}

func (s *Session) Delete() error {
	cmd := fmt.Sprintf("DELETE FROM %s WHERE sessionid = ?", s.TableName())
	_, err := DbConnection.Exec(cmd, s.SessionID)
	if err != nil {
		return err
	}
	return err
}

func GetSession(sessionID string) (*Session, error) {
	tableName := GetSessionTableName("session")
	cmd := fmt.Sprintf(`SELECT sessionid, username, time FROM %s WHERE sessionid = '%s'`,
		tableName, sessionID)
	row := DbConnection.QueryRow(cmd)
	var s Session
	err := row.Scan(&s.SessionID, &s.UserName, &s.Time)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return NewSession(s.SessionID, s.UserName, s.Time), nil
}

func GetRecentSessions(t time.Time) ([]Session, error) {
	tableName := GetSessionTableName("session")
	cmd := fmt.Sprintf(`SELECT sessionid, username, time FROM %s WHERE time > '%s'`,
		tableName, t)
	rows, err := DbConnection.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ss []Session
	for rows.Next() {
		var s Session
		err := rows.Scan(&s.SessionID, &s.UserName, &s.Time)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		ss = append(ss, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ss, nil
}

func GetOldSessions(t time.Time) ([]Session, error) {
	tableName := GetSessionTableName("session")
	cmd := fmt.Sprintf(`SELECT sessionid, username, time FROM %s WHERE time < '%s'`,
		tableName, t)
	rows, err := DbConnection.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ss []Session
	for rows.Next() {
		var s Session
		err := rows.Scan(&s.SessionID, &s.UserName, &s.Time)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		ss = append(ss, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ss, nil
}
