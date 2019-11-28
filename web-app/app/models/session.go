package models

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Session struct {
	SessionID string    `json:"session_id" bson:"session_id"`
	UserName  string    `json:"user_name" bson:"user_name"`
	Time      time.Time `json:"time" bson:"time"`
}

func NewSession(sessionId, userName string, timeDate time.Time) *Session {
	return &Session{
		SessionID: sessionId,
		UserName:  userName,
		Time:      timeDate,
	}
}

var sCollection *mongo.Collection

func init() {
	hrtfCollection = db.Collection("session")
}

func (s *Session) TableName() string {
	return GetSessionTableName("session")
}

func (s *Session) Create() error {
	//cmd := fmt.Sprintf("INSERT INTO %s (sessionid, username, time) VALUES (?, ?, ?)", s.TableName())
	//_, err := DbConnection.Exec(cmd, s.SessionID, s.UserName, s.Time.Format(tFormat))
	b, err := bson.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	_, err = sCollection.InsertOne(context.TODO(), b)
	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}
	return nil
}

func (s *Session) Save() error {
	//cmd := fmt.Sprintf("UPDATE %s SET username = ?, time = ? WHERE sessionid = ?", s.TableName())
	//_, err := DbConnection.Exec(cmd, s.UserName, s.Time.Format(tFormat), s.SessionID)

	filter := bson.D{{"session_id", s.SessionID}}
	b, err := bson.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	_, err = hrtfCollection.UpdateOne(context.TODO(), filter, b)
	if err != nil {
		return errors.Wrap(err, "failed to update data at UpdateOne()")
	}
	return nil
}

func (s *Session) Delete() error {
	//cmd := fmt.Sprintf("DELETE FROM %s WHERE sessionid = ?", s.TableName())
	//_, err := DbConnection.Exec(cmd, s.SessionID)

	filter := bson.D{{"session_id", s.SessionID}}
	_, err := hrtfCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}
	return nil
}

func GetSession(sessionID string) (*Session, error) {
	//tableName := GetSessionTableName("session")
	//cmd := fmt.Sprintf(`SELECT sessionid, username, time FROM %s WHERE sessionid = '%s'`,
	//	tableName, sessionID)
	//row := DbConnection.QueryRow(cmd)
	//var s Session
	//err := row.Scan(&s.SessionID, &s.UserName, &s.Time)

	filter := bson.D{{"session_id", sessionID}}

	var s Session
	err := sCollection.FindOne(context.TODO(), filter).Decode(&s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewSession(s.SessionID, s.UserName, s.Time), nil
}

func GetRecentSessions(t time.Time) ([]Session, error) {
	//tableName := GetSessionTableName("session")
	//cmd := fmt.Sprintf(`SELECT sessionid, username, time FROM %s WHERE time > '%s'`,
	//	tableName, t)
	//rows, err := DbConnection.Query(cmd)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//var ss []Session
	//for rows.Next() {
	//	var s Session
	//	err := rows.Scan(&s.SessionID, &s.UserName, &s.Time)
	//	if err != nil {
	//		log.Println(err)
	//		return nil, err
	//	}
	//	ss = append(ss, s)
	//}
	//if err := rows.Err(); err != nil {
	//	return nil, err
	//}
	//return ss, nil

	findOptions := options.Find()
	filter := bson.D{{"time", bson.D{{"$gt", t}}}}
	cur, err := sCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at Find()")
	}

	var ss []Session
	for cur.Next(context.TODO()) {
		var s Session
		if err = cur.Decode(&s); err != nil {
			return nil, errors.Wrap(err, "failed to decode data at Decode()")
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func GetOldSessions(t time.Time) ([]Session, error) {
	//tableName := GetSessionTableName("session")
	//cmd := fmt.Sprintf(`SELECT sessionid, username, time FROM %s WHERE time < '%s'`,
	//	tableName, t)
	//rows, err := DbConnection.Query(cmd)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//var ss []Session
	//for rows.Next() {
	//	var s Session
	//	err := rows.Scan(&s.SessionID, &s.UserName, &s.Time)
	//	if err != nil {
	//		log.Println(err)
	//		return nil, err
	//	}
	//	ss = append(ss, s)
	//}
	//if err := rows.Err(); err != nil {
	//	return nil, err
	//}
	//return ss, nil

	findOptions := options.Find()
	filter := bson.D{{"time", bson.D{{"$lt", t}}}}
	cur, err := sCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at Find()")
	}

	var ss []Session
	for cur.Next(context.TODO()) {
		var s Session
		if err = cur.Decode(&s); err != nil {
			return nil, errors.Wrap(err, "failed to decode data at Decode()")
		}
		ss = append(ss, s)
	}
	return ss, nil
}
