package main

import (
	"context"
	"gRPCpet/pkg/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewUserClient(conn)
	res, err := client.Create(context.Background(), &api.CreateRequest{
		Name:     "test",
		Email:    "test@mail.com",
		Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.String())
}
