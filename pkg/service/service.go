package service

import (
	"gRPCpet/pkg/entity"
	"gRPCpet/pkg/repository"
)

type Service struct {
	User User
}

type User interface {
	Create(user *entity.User) (uint64, error)
	GetAll() ([]entity.User, error)
	Delete(userId int) error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
