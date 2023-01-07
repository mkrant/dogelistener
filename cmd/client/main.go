package main

import (
	"github.com/mkrant/dogelistener/internal/iot"
	"os"
	"os/signal"
	"syscall"
)

const SockAddr = "/tmp/echo.sock"

func main() {
	srv := iot.NewServer(iot.ServerConfig{SocketAddr: SockAddr, ServerAddr: "localhost:8081"})

	go srv.Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
