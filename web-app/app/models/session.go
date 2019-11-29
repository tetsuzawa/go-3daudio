package models

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *Session) TableName() string {
	return GetTableName(tableNameSession)
}

func (s *Session) Create() error {
	sCollection := db.Collection(tableNameSession)

	b, err := bson.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	insertedID, err := sCollection.InsertOne(context.Background(), b)
	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}
	return nil
}

func (s *Session) Save() error {
	sCollection := db.Collection(tableNameSession)

	filter := bson.D{{"session_id", s.SessionID}}
	b, err := bson.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	_, err = sCollection.UpdateOne(context.TODO(), filter, b)
	if err != nil {
		return errors.Wrap(err, "failed to update data at UpdateOne()")
	}
	return nil
}

func (s *Session) Delete() error {
	sCollection := db.Collection(tableNameSession)

	filter := bson.D{{"session_id", s.SessionID}}
	_, err := sCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}
	return nil
}

func GetSession(sessionID string) (*Session, error) {
	sCollection := db.Collection(tableNameSession)

	filter := bson.D{{"session_id", sessionID}}

	var s Session
	err := sCollection.FindOne(context.Background(), filter).Decode(&s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	fmt.Println("got session:", s)
	return NewSession(s.SessionID, s.UserName, s.Time), nil
}

func GetRecentSessions(t time.Time) ([]Session, error) {
	sCollection := db.Collection(tableNameSession)

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
	sCollection := db.Collection(tableNameSession)

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
