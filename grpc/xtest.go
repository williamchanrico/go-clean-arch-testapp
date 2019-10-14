package grpc

import (
	"context"
	"os"

	"github.com/williamchanrico/xtest/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Server for the Xtest gRPC API
type Server struct{}

// Xtest grpc for testing
func (s *Server) Xtest(ctx context.Context, req *XtestRequest) (*XtestResponse, error) {
	log.Debugf("Handling GRPC Xtest request [%v] with context %v\n", req, ctx)
	hostname, err := os.Hostname()
	if err != nil {
		log.Errorf("Unable to get hostname: %v", err.Error())
		hostname = ""
	}
	grpc.SendHeader(ctx, metadata.Pairs("hostname", hostname))
	return &XtestResponse{Content: req.Content}, nil
}
