package handlers

import "database/sql"

type Handler struct {
	Db *sql.DB
}

func CreateHandler() *Handler {
	handler := &Handler{}
	return handler
}
