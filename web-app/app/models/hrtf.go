package models

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type HRTF struct {
	ID           string `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	Path         string `json:"path" bson:"path"`
	DatabaseName string `json:"database_name" bson:"database_name"`
}

func NewHRTF(id, name, path, databaseName string) *HRTF {
	return &HRTF{
		ID:           id,
		Name:         name,
		Path:         path,
		DatabaseName: databaseName,
	}
}

func (h *HRTF) TableName() string {
	return GetTableName(tableNameHRTFData)
}

func (h *HRTF) Create() error {
	hrtfCollection := db.Collection(h.TableName())

	b, err := bson.Marshal(h)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	_, err = hrtfCollection.InsertOne(context.TODO(), b)
	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}
	return nil
}

func (h *HRTF) Save() error {
	hrtfCollection := db.Collection(h.TableName())

	filter := bson.D{{"id", h.ID}}
	b, err := bson.Marshal(h)
	if err != nil {
		return errors.Wrap(err, "failed to encode at bson.Marshal()")
	}
	_, err = hrtfCollection.UpdateOne(context.TODO(), filter, b)
	if err != nil {
		return errors.Wrap(err, "failed to insert data at InsertOne()")
	}
	return nil
}

func GetHRTF(id string) (*HRTF, error) {
	hrtfCollection := db.Collection(GetTableName(tableNameHRTFData))

	filter := bson.D{{"id", id}}

	var hrtf HRTF
	err := hrtfCollection.FindOne(context.TODO(), filter).Decode(&hrtf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewHRTF(hrtf.ID, hrtf.Name, hrtf.Path, hrtf.DatabaseName), nil
}

func GetHRTFFromName(name string) (*HRTF, error) {
	hrtfCollection := db.Collection(GetTableName(tableNameHRTFData))

	filter := bson.D{{"name", name}}

	var hrtf HRTF
	err := hrtfCollection.FindOne(context.TODO(), filter).Decode(&hrtf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find data at FindOne()")
	}
	return NewHRTF(hrtf.ID, hrtf.Name, hrtf.Path, hrtf.DatabaseName), nil
}
