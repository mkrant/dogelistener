package iot

import "log"

type Server struct {
	ss  *SocketServer
	api *API
}

type ServerConfig struct {
	SocketAddr string
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{ss: NewSocketServer(cfg.SocketAddr)}
}

func (s *Server) Run() {
	log.Printf("Running socket server using file %s", s.ss.sockAddr)

	go s.ss.Run()
}
