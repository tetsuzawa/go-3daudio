package models

import (
	"fmt"
	"log"
)

type HRTF struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Age       uint    `json:"age"`
	Azimuth   float64 `json:"azimuth"`
	Elevation float64 `json:"elevation"`
	Data      float64 `json:"data"`
}

func NewHRTF(id int, name string, age uint, azimuth, elevation, data float64) *HRTF {
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
	return GetHRTFTableName("hrtf")
}

func (h *HRTF) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (id, name, age, azimuth, elevation, data) VALUES (?, ?, ?, ?, ?, ?)", h.TableName())
	_, err := DbConnection.Exec(cmd, h.ID, h.Name, h.Age, h.Azimuth, h.Elevation, h.Data)
	if err != nil {
		return err
	}
	return err
}

func (h *HRTF) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET name = ?, age = ?, azimuth = ?, elevation = ?, data = ? WHERE time = ?", h.TableName())
	_, err := DbConnection.Exec(cmd, h.Name, h.Age, h.Azimuth, h.Elevation, h.Data, h.Name)
	if err != nil {
		return err
	}
	return err
}

func GetHRTF(id string) (*HRTF, error){
	//tableName := GetHRTFTableName(string(id))
	tableName := GetHRTFTableName("hrtf")
	cmd := fmt.Sprintf(`SELECT id, name, age, azimuth, elevation, data FROM %s WHERE id = '%s'`,
		tableName, id)
	row := DbConnection.QueryRow(cmd)
	var hrtf HRTF
	err := row.Scan(&hrtf.ID, &hrtf.Name, &hrtf.Age, &hrtf.Azimuth, &hrtf.Elevation, &hrtf.Data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return NewHRTF(hrtf.ID, hrtf.Name, hrtf.Age, hrtf.Azimuth, hrtf.Elevation, hrtf.Data), nil
}
