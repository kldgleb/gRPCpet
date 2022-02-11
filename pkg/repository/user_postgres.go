package repository

import (
	"database/sql"
	"gRPCpet/pkg/entity"
	"log"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db}
}

func (r *UserPostgres) Create(user *entity.User) (uint64, error) {
	log.Println(user)
	return 1, nil
}

func (r *UserPostgres) GetAll() ([]entity.User, error) {
	var users []entity.User
	users = append(users, entity.User{
		Id:       1,
		Name:     "User",
		Email:    "test@mail.com",
		Password: "password",
	})
	return users, nil
}

func (r *UserPostgres) Delete(userId int) error {
	log.Println(userId)
	return nil
}
