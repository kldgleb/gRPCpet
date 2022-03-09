package grpc

import (
	"gRPCpet/pkg/api"
	"gRPCpet/transport/grpc/handler"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	srv          *grpc.Server
	Dependencies Dependencies
}

type Dependencies struct {
	UserHandler *handler.UserHandler
}

type ServerConfig struct {
	Host string
	Port string
}

func NewServer(d Dependencies) *Server {
	return &Server{
		srv:          grpc.NewServer(),
		Dependencies: d,
	}
}

func (s *Server) ListenAndServe(cfg ServerConfig) error {
	api.RegisterUserServer(s.srv, s.Dependencies.UserHandler)

	l, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		return err
	}
	if err = s.srv.Serve(l); err != nil {
		return err
	}
	return nil
}
