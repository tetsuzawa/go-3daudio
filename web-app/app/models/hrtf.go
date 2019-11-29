package models

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type HRTF struct {
	ID        string  `json:"id" bson:"id"`
	Name      string  `json:"name" bson:"name"`
	Age       uint    `json:"age" bson:"age"`
	Azimuth   float64 `json:"azimuth" bson:"azimuth"`
	Elevation float64 `json:"elevation" bson:"elevation"`
	Data      float64 `json:"data" bson:"data"`
}

func NewHRTF(id string, name string, age uint, azimuth, elevation, data float64) *HRTF {
	return &HRTF{
		ID:        id,
		Name:      name,
		Age:       age,
		Azimuth:   azimuth,
		Elevation: elevation,
		Data:      data,
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
	return NewHRTF(hrtf.ID, hrtf.Name, hrtf.Age, hrtf.Azimuth, hrtf.Elevation, hrtf.Data), nil
}
