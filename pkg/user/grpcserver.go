package user

import (
	"context"
	"gRPCpet/pkg/api"
	"gRPCpet/pkg/entity"
	"gRPCpet/pkg/service"
	"log"
)

type GRPCServer struct {
	Service *service.Service
	api.UnimplementedUserServer
}

func (s *GRPCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	userId, err := s.Service.User.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	return &api.CreateResponse{Id: userId}, nil
}
