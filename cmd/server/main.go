package main

import (
	"gRPCpet/pkg/api"
	"gRPCpet/pkg/repository"
	"gRPCpet/pkg/service"
	"gRPCpet/pkg/user"
	"github.com/joho/godotenv"
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

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  viper.GetString("db.sslMode"),
	})
	if err != nil {
		log.Fatal("Connect to db err: ", err)
	}

	rdb, err := repository.NewRedisClient(repository.RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: viper.GetString("rdb.password"),
		DB:       viper.GetInt("rdb.db"),
	})
	if err != nil {
		log.Fatal("Connect to rdb err: ", err)
	}

	repos := repository.NewRepository(db, rdb)
	services := service.NewService(repos)

	s := grpc.NewServer()
	userServer := &user.GRPCServer{
		Service: services,
	}
	api.RegisterUserServer(s, userServer)
	l, err := net.Listen("tcp", os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT"))
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
