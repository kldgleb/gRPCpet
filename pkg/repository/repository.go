package repository

import (
	"database/sql"
	"gRPCpet/pkg/entity"
	"github.com/go-redis/redis/v8"
)

type Repository struct {
	User
}

type User interface {
	Create(user *entity.User) (uint64, error)
	GetAll() ([]entity.User, error)
	Delete(userId uint64) error
	CacheUsers([]entity.User)
	GetCachedUsers() ([]entity.User, error)
	HasCachedUsers() bool
}

func NewRepository(db *sql.DB, rdb *redis.Client) *Repository {
	return &Repository{
		User: NewUserRepository(db, rdb),
	}
}
