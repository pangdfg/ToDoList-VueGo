package grpc

import (
	"google.golang.org/grpc"
)

type Server struct {
	engine *grpc.Server
}

func NewServer() *Server {
	return &Server{
		engine: grpc.NewServer(
		),
	}
}	