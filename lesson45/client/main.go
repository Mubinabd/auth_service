package main

import (
	"context"
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cs := pb.NewCoffeeServiceClient(conn)

	res, err := cs.BuyingCoffee(context.Background(), &pb.BuyCoffee{Name: "Latte", IsPaid: true})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", res)
}
