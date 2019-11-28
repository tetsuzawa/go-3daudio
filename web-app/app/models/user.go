package models

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

var userCollection *mongo.Collection

func init() {
	userCollection = db.Collection("user")
}

func (u *User) TableName() string {
	return GetUserTableName("user")
}

func (u *User) Create() error {
	//cmd := fmt.Sprintf("INSERT INTO %s (id, username, password, firstname, lastname, role) VALUES (?, ?, ?, ?, ?, ?)", u.TableName())
	//_, err := DbConnection.Exec(cmd, u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role)
	//if err != nil {
	//	return err
	//}

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
	//cmd := fmt.Sprintf("UPDATE %s SET username = ?, password = ?, firstname = ?, lastname = ?, role = ? WHERE id = ?", u.TableName())
	//_, err := DbConnection.Exec(cmd, u.UserName, u.Password, u.FirstName, u.LastName, u.Role, u.ID)
	//if err != nil {
	//	return err
	//}
	//return err

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
	//tableName := GetUserTableName("user")
	//cmd := fmt.Sprintf(`SELECT id, username, password, firstname, lastname, role FROM %s WHERE id = '%s'`,
	//	tableName, id)
	//row := DbConnection.QueryRow(cmd)
	//var u User
	//err := row.Scan(&u.ID, &u.UserName, &u.Password, &u.FirstName, &u.LastName, &u.Role)
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}
	//return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil

	filter := bson.D{{"id", id}}

	var u User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}

func GetUserByUserName(un string) (*User, error) {
	//tableName := GetUserTableName("user")
	//cmd := fmt.Sprintf(`SELECT id, username, password, firstname, lastname, role FROM %s WHERE username = '%s'`,
	//	tableName, un)
	//row := DbConnection.QueryRow(cmd)
	//var u User
	//err := row.Scan(&u.ID, &u.UserName, &u.Password, &u.FirstName, &u.LastName, &u.Role)
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}
	//return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil

	filter := bson.D{{"user_name", un}}

	var u User
	err := sCollection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}
