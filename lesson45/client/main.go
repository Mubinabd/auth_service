package main

import (
	"context"
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	pbcr "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cs := pb.NewCoffeeServiceClient(conn)
	cr := pbcr.NewCourierServiceClient(conn)

	res, err := cs.BuyingCoffee(context.Background(), &pb.BuyCoffee{Name: "Latte", IsPaid: true})
	if err != nil {
		panic(err)
	}
	order := &pbcr.TakeOrder{
		Name:    "Order 1",
		IsPaid:  true,
		Address: "Qatortol 9",
	}
	delResp, err := cr.Deliver(context.Background(), order)
	if err != nil {
		log.Fatalf("error delResp err: %s", err.Error())
	}

	fmt.Printf("%v\n", res)

	fmt.Printf("%v\n", delResp)
}
