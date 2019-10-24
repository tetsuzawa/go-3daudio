package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tetsuzawa/go-3daudio/config"
)

const (
	tableNameHRTFData = "hrtf"
	tableNameUserData = "user"
	tableNameSession  = "session"
)

const tFormat = "2006-01-02 15:04:05"

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
		id TEXT NOT NULL,
		name TEXT,
		age INT,
		azimuth FLOAT, 
		elevation FLOAT, 
		data FLOAT,
		PRIMARY KEY(id(128)))`, tableNameHRTFData)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	cmd = fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
		id TEXT NOT NULL,
		username TEXT,
		password TEXT,
		firstname TEXT, 
		lastname TEXT,
		role TEXT,
		PRIMARY KEY(id(128)))`, tableNameUserData)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	cmd = fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
		sessionid TEXT NOT NULL,
		username TEXT,
		time DATETIME,
		PRIMARY KEY(sessionid(128)))`, tableNameSession)
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
