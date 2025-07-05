package server

import (
	"context"
	"github.com/bitmeme-taxi/bitmemed/cmd/gorwallet/daemon/pb"
	"github.com/bitmeme-taxi/bitmemed/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
