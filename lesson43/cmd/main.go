package main

import (
	"fmt"
	"github.com/husanmusa/NT_Golang_10/lesson43/api"
	"github.com/husanmusa/NT_Golang_10/lesson43/api/handlers"
	"github.com/husanmusa/NT_Golang_10/lesson43/config"
	"github.com/husanmusa/NT_Golang_10/lesson43/service"
	"github.com/husanmusa/NT_Golang_10/lesson43/storage"
	"log"
	"regexp"
)

func main() {
	b, err := regexp.Match("\\+\\((?:998\\))([0-9]{2})[-]([0-9]{3})[-](\\d{2})[-](\\d{2})",
		[]byte("+(998)92-435-22-21"))
	fmt.Println(b)
	return
	cfg := config.Load()
	db, err := storage.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error while connect db, err: %s", err.Error())
	}

	services := service.InitServices(db)

	handler := handlers.NewHTTPHandler(services)

	engine := api.NewGin(handler)

	err = engine.Run(cfg.HTTPPort)
	if err != nil {
		log.Fatalf("error while run engine, err: %s", err.Error())
	}
}
