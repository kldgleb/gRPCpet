package repository

import (
	"database/sql"
	"gRPCpet/pkg/entity"
)

type Repository struct {
	User
}

type User interface {
	Create(user *entity.User) (uint64, error)
	GetAll() ([]entity.User, error)
	Delete(userId int) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
