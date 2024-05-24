package handler

import (
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/postgres"
	"github.com/jackc/pgx/v5"
)

type HandlerStruct struct {
	User   postgres.UserDB
	Task   postgres.TaskDB
	Config config.Config
}

func NewHandler(db *pgx.Conn, config config.Config) *HandlerStruct {
	return &HandlerStruct{
		User:   *postgres.NewUser(db),
		Task:   *postgres.NewTask(db),
		Config: config.Config,
	}
}
