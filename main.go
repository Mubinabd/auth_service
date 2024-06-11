package main

import (
	"github.com/Mubinabd/auth_service/api"
	h "github.com/Mubinabd/auth_service/api/handler"
	"github.com/Mubinabd/auth_service/service"
	"github.com/Mubinabd/auth_service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	userService := service.NewUserService(db)
	h := h.NewHandler(userService)

	router := api.NewGin(h)
	router.Run(":8080")
}
