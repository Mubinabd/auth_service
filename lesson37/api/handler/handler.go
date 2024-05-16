package handler

import "database/sql"

type Handler struct {
	Db *sql.DB
}

// New ...
func New(db *sql.DB) *Handler {
	return &Handler{db}
}
