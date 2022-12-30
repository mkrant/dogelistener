package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type API struct {
	UnimplementedDogeServerServer
}

func (s *API) Connect(stream DogeServer_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
