package main

import (
	"fmt"
	"log"
	"net"
	"server/config"
	pbu "server/genproto/user"
	"server/pkg/db"
	"server/pkg/logger"
	"server/service"
	grpcclient "server/service/grpc-client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg := config.Load()

	connDb, err := db.ConnectToDB(cfg)
	if err != nil {
		fmt.Println("failed connect database", err)
	}

	grpcClient, err := grpcclient.NewServiceMeneger(cfg)
	if err != nil {
		fmt.Println("failed while grpc client", err.Error())
	}

	userService := service.NewUserServer(connDb, grpcClient)

	listener, err := net.Listen("tcp", cfg.UserServicePort)
	if err != nil {
		log.Fatal("failed while listening port: %v", logger.Error(err))
	}
	
	grpS := grpc.NewServer()
	reflection.Register(grpS)

	pbu.RegisterUserServiceServer(grpS, userService)
	log.Println("main: server running", logger.String("port", cfg.UserServicePort))

	err = grpS.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
