package main

import "github.com/mkrant/dogelistener/internal/server"

func main() {
	srv := server.NewServer()
	srv.Start()
}
