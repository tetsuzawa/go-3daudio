package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tetsuzawa/go-3daudio/config"
	"log"
)

const (
	tableNameHRTFData = "hrtf"
	tableNameUserData = "user"
	tableNameSession  = "session"
)

var DbConnection *sql.DB

func GetHRTFTableName(name string) string {
	return fmt.Sprintf("%s", name)
}

func GetUserTableName(name string) string {
	return fmt.Sprintf("%s", name)
}

func GetSessionTableName(name string) string {
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
		id STRING PRIMARY KEY NOT NULL,
		name STRING,
		age INT,
		azimuth FLOAT, 
		elevation FLOAT, 
		data FLOAT)`, tableNameHRTFData)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	cmd = fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
		id STRING PRIMARY KEY NOT NULL,
		username STRING,
		password STRING,
		firstname STRING, 
		lastname STRING,
		role STRING)`, tableNameUserData)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	cmd = fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
		sessionid STRING PRIMARY KEY NOT NULL,
		username STRING)`, tableNameSession)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
