package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tetsuzawa/go-3daudio/config"
	"log"
)

const tableNameHRTFData = "hrtf_data"

var DbConnection *sql.DB

func GetHRTFTableName(name string) string {
	return fmt.Sprintf("%s", name)
}

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
		time DATETIME PRIMARY KEY NOT NULL,
		id int,
		name STRING,
		azimuth FLOAT, 
		elevation FLOAT, 
		data FLOAT)`, tableNameHRTFData)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
