package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/literaen/simple_project/pkg/grpc/interceptor"

	"time"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewServer(timeout time.Duration) *Server {
	return &Server{
		server: grpc.NewServer(
			grpc.ConnectionTimeout(timeout),
			grpc.ChainUnaryInterceptor(
				interceptor.Interceptor,
			),
		),
	}
}

func (s *Server) Start(ctx context.Context, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s.listener = lis

	log.Printf("Starting gRPC server on address %s", address)
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (s *Server) GetServer() *grpc.Server {
	return s.server
}

func (s *Server) Stop() error {
	if s.server != nil {
		s.server.Stop()
	}

	return nil
}
