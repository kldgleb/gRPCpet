package main

import (
	"context"
	"gRPCpet/pkg/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewUserClient(conn)
	//res, err := client.Create(context.Background(), &api.CreateRequest{
	//	Name:     "test",
	//	Email:    "test@mail.com",
	//	Password: "password",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Println(res.String())

	usersRes, err := client.Users(context.Background(), &api.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(usersRes.String())

	delRes, err := client.Delete(context.Background(), &api.DeleteRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(delRes.String())

}
