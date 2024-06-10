package main

import (
	"log"
	"net"
	"github.com/Mubinabd/auth_service/config"
	pb "github.com/Mubinabd/auth_service/genproto"
	"github.com/Mubinabd/auth_service/service"
	"github.com/Mubinabd/auth_service/storage/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()
	db, err := postgres.ConnectDb(cfg)
	if err != nil {
		log.Println("error while connecting to postgres: ", err)
	}

	liss, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.NewUserService(db))

	reflection.Register(s)
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}