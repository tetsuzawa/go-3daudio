package models

import (
	"fmt"
)

type HRTF struct {
	Name      string  `json:"name"`
	Age       string  `json:"age"`
	Azimuth   float64 `json:"azimuth"`
	Elevation float64 `json:"elevation"`
	Data      float64 `json:"data"`
}

func NewHRTF(name, age string, azimuth, elevation, data float64) *HRTF {
	return &HRTF{
		Name:      name,
		Age:       age,
		Azimuth:   azimuth,
		Elevation: elevation,
		Data:      data,
	}
}

func (h *HRTF) TableName() string {
	return GetHRTFTableName(h.Name)
}

func (h *HRTF) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s (name, age, azimuth, elevation, data) VALUES (?, ?, ?, ?, ?)", h.TableName())
	_, err := DbConnection.Exec(cmd, h.Name, h.Age, h.Azimuth, h.Elevation, h.Data)
	if err != nil {
		return err
	}
	return err
}

func (h *HRTF) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET age = ?, azimuth = ?, elevation = ?, data = ? WHERE time = ?", h.TableName())
	_, err := DbConnection.Exec(cmd, h.Age, h.Azimuth, h.Elevation, h.Data, h.Name)
	if err != nil {
		return err
	}
	return err
}

func GetHRTF(name string) *HRTF {
	tableName := GetHRTFTableName(name)
	cmd := fmt.Sprintf(`SELECT name, age, azimuth, elevation, data FROM %s WHERE name = '%s'`,
		tableName, name)
	row := DbConnection.QueryRow(cmd)
	var hrtf HRTF
	err := row.Scan(&hrtf.Name, &hrtf.Age, &hrtf.Azimuth, &hrtf.Elevation, &hrtf.Data)
	if err != nil {
		return nil
	}
	return NewHRTF(hrtf.Name, hrtf.Age, hrtf.Azimuth, hrtf.Elevation, hrtf.Data)
}
