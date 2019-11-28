package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/tetsuzawa/go-3daudio/web-app/config"
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
	err = godotenv.Load() //Load env.file
	if err != nil {
		log.Fatalln(err)
	}
	dbName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		config.Cfg.DB.Host,
		config.Cfg.DB.Port,
		config.Cfg.DB.Name,
		config.Cfg.DB.ETC,
	)
	//dbName := fmt.Sprintf(config.Cfg.DB.Name, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	DbConnection, err = sql.Open(config.Cfg.DB.Driver, dbName)
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
