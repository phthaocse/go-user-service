package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/phthaocse/user-service-go/config"
)

func SetUp(config *config.Config) (*sql.DB, error) {
	dbConfig := &mysql.Config{
		User: config.DbUser,
		Addr: config.DbAddr,
		DBName: config.DbName,
		Passwd: config.DbPassword,
		AllowNativePasswords: true,
	}
	var err error
	var db *sql.DB
	dbUri := dbConfig.FormatDSN()
	db, err = sql.Open(config.DbDriver, dbUri)
	if err != nil {
		fmt.Println("Unable to connect to db with err: ", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Unable to connect to db with err: ", err)
		return nil, err
	}
	fmt.Println("Connect to db successfully")
	return db, nil
}
