package grpc

import (
	"context"

	xtest "github.com/williamchanrico/xtest/grpc"
	"github.com/williamchanrico/xtest/listener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server for grpc service
type Server struct {
	address string

	server *grpc.Server
}

// New grpc server
func New(address string) *Server {
	return &Server{
		address: address,
	}
}

// Run grpc server
func (s *Server) Run() error {
	l, err := listener.Listen(s.address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	xtest.RegisterXtestServer(grpcServer, &xtest.Server{})
	reflection.Register(grpcServer)

	s.server = grpcServer

	return grpcServer.Serve(l)
}

// Shutdown GRPC server
func (s *Server) Shutdown(ctx context.Context) {
	if s.server == nil {
		return
	}
	s.server.GracefulStop()
}
