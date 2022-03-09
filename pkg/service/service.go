package service

import (
	"gRPCpet/pkg/repository"
	"gRPCpet/transport/grpc/handler"
)

type Service struct {
	User handler.UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
