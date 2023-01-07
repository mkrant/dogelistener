package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/mkrant/dogelistener/internal/server/api"
	"github.com/mkrant/dogelistener/internal/server/sessions"
	"io"
	"log"
)

type API struct {
	api.UnimplementedDogeServerServer
	sessManager *sessions.Manager
	counter     int
}

func NewAPI(sessManager *sessions.Manager) *API {
	return &API{
		sessManager: sessManager,
	}
}

func (a *API) Connect(stream api.DogeServer_ConnectServer) error {
	sess, err := a.sessManager.StartSession(stream)
	if err != nil {
		return fmt.Errorf("start session: %w", err)
	}

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
			return nil
		}
		if err != nil {
			log.Printf("cannot receive %v", err)
			return err
		}

		switch req := resp.Type.(type) {
		case *api.Request_Ping:
			log.Println("Got ping")
			stream.Send(&api.Response{Type: &api.Response_Pong{Pong: &api.Pong{}}})
		case *api.Request_RunData:
			log.Println("Got run data", req)
		}

	}
}
