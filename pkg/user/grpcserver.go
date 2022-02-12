package user

import (
	"context"
	"gRPCpet/pkg/api"
	"gRPCpet/pkg/entity"
	"gRPCpet/pkg/service"
)

type GRPCServer struct {
	Service *service.Service
	api.UnimplementedUserServer
}

func (s *GRPCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.UserResponse, error) {
	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	userId, err := s.Service.User.Create(user)
	return &api.UserResponse{
		Id:       userId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, err
}

func (s *GRPCServer) Users(context.Context, *api.Empty) (*api.UsersResponse, error) {
	users, err := s.Service.User.GetAll()
	response := &api.UsersResponse{Users: []*api.UserResponse{}}
	for i := range users {
		response.Users = append(response.Users, &api.UserResponse{
			Id:       users[i].Id,
			Name:     users[i].Name,
			Email:    users[i].Email,
			Password: users[i].Password,
		})
	}
	return response, err
}

func (s *GRPCServer) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {
	userId := req.GetId()
	err := s.Service.User.Delete(userId)
	response := &api.DeleteResponse{
		Message: "OK",
	}
	if err != nil {
		response.Message = "Err"
	}
	return response, err
}
