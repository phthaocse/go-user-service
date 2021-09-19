package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/phthaocse/user-service-go/config"
	"log"
)

func SetUp(config *config.Config, log *log.Logger) (*sql.DB, error) {
	dbConfig := &mysql.Config{
		User:                 config.DbUser,
		Addr:                 config.DbAddr,
		DBName:               config.DbName,
		Passwd:               config.DbPassword,
		AllowNativePasswords: true,
	}
	var err error
	var db *sql.DB
	dbUri := dbConfig.FormatDSN()
	db, err = sql.Open(config.DbDriver, dbUri)
	if err != nil {
		log.Println("Unable to connect to db with err: ", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Unable to connect to db with err: ", err)
		return nil, err
	}
	log.Println("Connect to db successfully")
	return db, nil
}
