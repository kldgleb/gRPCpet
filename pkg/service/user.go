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
	userId, err := s.repo.User.Create(user)
	if err == nil {
		s.repo.User.FlushCachedUsers()
	}
	return userId, err
}

func (s *UserService) GetAll() ([]entity.User, error) {

	hasCache := s.repo.User.HasCachedUsers()
	if hasCache {
		users, err := s.repo.User.GetCachedUsers()
		return users, err
	}

	users, err := s.repo.User.GetAll()
	s.repo.User.CacheUsers(users)
	return users, err
}

func (s *UserService) Delete(userId uint64) error {
	return s.repo.User.Delete(userId)
}
