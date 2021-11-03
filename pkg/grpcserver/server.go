// Package grpcserver implements gRPC server.
package grpcserver

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	_defaultAddr = ":50051"
)

var RpcServer *grpc.Server

// Server -.
type Server struct {
	server *grpc.Server
	notify chan error
	Addr   string
}

// New -.
func New(opts ...Option) *Server {
	grpcServer := grpc.NewServer()
	// 注册 grpcurl 所需的 reflection 服务
	reflection.Register(grpcServer)

	RpcServer = grpcServer

	s := &Server{
		server: grpcServer,
		notify: make(chan error, 1),
		Addr:   _defaultAddr,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		lis, err := net.Listen("tcp", s.Addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s.notify <- s.server.Serve(lis)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
