package service

import (
	"gRPCpet/pkg/entity"
	"gRPCpet/pkg/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Create(user *entity.User) (uint64, error) {
	return s.repo.User.Create(user)
}

func (s *UserService) GetAll() ([]entity.User, error) {
	return s.repo.User.GetAll()
}

func (s *UserService) Delete(userId uint64) error {
	return s.repo.User.Delete(userId)
}
