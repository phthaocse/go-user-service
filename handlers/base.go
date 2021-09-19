package handlers

import (
	"database/sql"
	"log"
	"os"
)

type Handler struct {
	Db  *sql.DB
	Log *log.Logger
}

func CreateHandler() *Handler {
	handler := &Handler{Log: log.New(os.Stdout, "", log.LstdFlags)}
	return handler
}
