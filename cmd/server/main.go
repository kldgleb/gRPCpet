package main

import (
	"gRPCpet/pkg/api"
	"gRPCpet/pkg/repository"
	"gRPCpet/pkg/service"
	"gRPCpet/pkg/user"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while reading config, %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  viper.GetString("db.sslMode"),
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	s := grpc.NewServer()
	userServer := &user.GRPCServer{
		Service: services,
	}
	api.RegisterUserServer(s, userServer)

	l, err := net.Listen("tcp", os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
