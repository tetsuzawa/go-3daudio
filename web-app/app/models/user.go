package models

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	userpb "github.com/tetsuzawa/go-3daudio/web-app/proto/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type User struct {
	ID        string `json:"id" bson:"id"`
	UserName  string `json:"user_name" bson:"user_name"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Role      string `json:"role" bson:"role"`
}

type UserServicer struct{}

func NewUser(id, username, password, firstname, lastname, role string) *userpb.User {
	return &userpb.User{
		Id:        id,
		UserName:  username,
		Password:  password,
		FirstName: firstname,
		LastName:  lastname,
		Role:      role,
	}
}

func (u *UserServicer) TableName() string {
	return GetTableName(tableNameUserData)
}

func (u *UserServicer) Create(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserRes, error) {
	userCollection := db.Collection(u.TableName())

	user := req.GetUser()

	data := userpb.User{
		//ID:        user.Id,  //empty, Mongodb generates a unique object ID
		UserName:  user.GetUserName(),
		Password:  user.GetPassword(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Role:      user.GetRole(),
	}

	//b, err := bson.Marshal(data)
	//if err != nil {
	//	log.Printf("failed to encode at bson.Marshal(): %v\n", err)
	//	return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	//}
	result, err := userCollection.InsertOne(context.TODO(), data)
	if err != nil {
		// return internal gRPC error to be handled later
		//return errors.Wrap(err, "failed to insert data at InsertOne()")
		log.Printf("failed to insert document at InsertOne: %v\n", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	// add the id to user, first cast the "generic type" (go doesn't have real generics yet) to an Object ID.
	old := result.InsertedID.(primitive.ObjectID)
	//	// Convert the object id to it's string counterpart
	user.Id = old.Hex()
	return &userpb.CreateUserRes{User: user}, nil
}

func (u *UserServicer) Read(ctx context.Context, req *userpb.ReadUserReq) (*userpb.ReadUserRes, error) {
	userCollection := db.Collection(u.TableName())

	id := req.GetId()

	filter := bson.D{{"id", id}}
	//b, err := bson.Marshal(u)
	//if err != nil {
	//	return errors.Wrap(err, "failed to encode at bson.Marshal()")
	//}

	var user *userpb.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		log.Printf("failed to insert document at FindOne: %v\n", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}
	return &userpb.ReadUserRes{User: user,}, nil
}

//_, err = userCollection.UpdateOne(context.TODO(), filter, id)
//if err != nil {
//return errors.Wrap(err, "failed to insert data at InsertOne()")
//}
//return nil
//
func GetUser(id string) (*UserServicer, error) {
	userCollection := db.Collection(GetTableName(tableNameUserData))

	filter := bson.D{{"id", id}}

	var u UserServicer
	err := userCollection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}

func GetUserByUserName(un string) (*UserServicer, error) {
	userCollection := db.Collection(GetTableName(tableNameUserData))

	filter := bson.D{{"user_name", un}}

	var u UserServicer
	err := userCollection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewUser(u.ID, u.UserName, u.Password, u.FirstName, u.LastName, u.Role), nil
}
