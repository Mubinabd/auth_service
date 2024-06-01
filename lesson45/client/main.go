package main

import (
	"context"
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	pbcr "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
	"github.com/husanmusa/NT_Golang_10/lesson45/grpc/client"
	"log"
)

func main() {
	clients, err := client.NewGrpcClients()
	if err != nil {
		log.Fatalln(err)
	}

	res, err := clients.CoffeeService.BuyingCoffee(context.Background(), &pb.BuyCoffee{Name: "Latte", IsPaid: true})
	if err != nil {
		panic(err)
	}
	order := &pbcr.TakeOrder{
		Name:    "Order 1",
		IsPaid:  true,
		Address: "Qatortol 9",
	}
	delResp, err := clients.CourierService.Deliver(context.Background(), order)
	if err != nil {
		log.Fatalf("error delResp err: %s", err.Error())
	}

	fmt.Printf("%v\n", res)

	fmt.Printf("%v\n", delResp)
}
