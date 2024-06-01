package main

import (
	pbcr "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
	"github.com/husanmusa/NT_Golang_10/lesson45/grpc/client"
	"github.com/husanmusa/NT_Golang_10/lesson45/service"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	services, err := client.NewGrpcClients()
	if err != nil {
		log.Fatalln(err)
	}

	liss, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	//pb.RegisterCoffeeServiceServer(s, service.NewCoffeeService(db, rdb))

	pbcr.RegisterCourierServiceServer(s, service.NewCourierService(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
