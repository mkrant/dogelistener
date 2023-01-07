package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mkrant/dogelistener/internal/server/api"
	"github.com/mkrant/dogelistener/internal/server/sessions"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type Server struct {
	grpcServer *grpc.Server
	api        *API
}

func NewServer() *Server {
	return &Server{
		api: NewAPI(sessions.NewManager()),
	}
}

func (s *Server) Start() {
	log.Println("Starting http server on port 8080")
	log.Println("Starting grpc server on port 8081")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterDogeServerServer(srv, s.api)
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	router := httprouter.New()
	s.api.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
