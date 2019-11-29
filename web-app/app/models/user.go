package models

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID        string `json:"id" bson:"id"`
	UserName  string `json:"user_name" bson:"user_name"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Role      string `json:"role" bson:"role"`
}

func NewUser(id, username, password, firstname, lastname, role string) *User {
	return &User{
		ID:        id,
		UserName:  username,
		Password:  password,
		FirstName: firstname,
		LastName:  lastname,
		Role:      role,
	}
}

func (u *User) TableName() string {
	return GetTableName(tableNameUserData)
}

func (u *User) Create() error {
	userCollection := db.Collection(u.TableName())

	b, err := bson.Marshal(u)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	_, err = userCollection.InsertOne(context.TODO(), b)
	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}

	return nil
}

func (u *User) Save() error {
	userCollection := db.Collection(u.TableName())

	filter := bson.D{{"id", u.ID}}
	b, err := bson.Marshal(u)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	_, err = userCollection.UpdateOne(context.TODO(), filter, b)
	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}
	return nil
}

func GetUser(id string) (*User, error) {
	userCollection := db.Collection(GetTableName(tableNameUserData))

	filter := bson.D{{"id", id}}

	var u User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}

func GetUserByUserName(un string) (*User, error) {
	userCollection := db.Collection(GetTableName(tableNameUserData))

	filter := bson.D{{"user_name", un}}

	var u User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}
