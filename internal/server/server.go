package server

import (
	"github.com/mkrant/dogelistener/internal/server/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	api        *api.API
}

func NewServer() *Server {
	return &Server{
		api: &api.API{},
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterDogeServerServer(srv, s.api)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
