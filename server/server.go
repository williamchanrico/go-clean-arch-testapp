package server

import (
	"context"

	xtestHTTP "github.com/williamchanrico/xtest/server/http"
)

// Server struct
type Server struct {
	HTTPAddress string

	stopChan chan context.Context
}

// Run xtest servers
func (s *Server) Run() error {
	httpServer := xtestHTTP.New(s.HTTPAddress)
	errChan := make(chan error)

	// Start HTTP server
	go func(httpServer *xtestHTTP.Server) {
		err := httpServer.Run()
		errChan <- err
	}(httpServer)

	select {
	case ctx := <-s.stopChan:
		return httpServer.Shutdown(ctx)
	case err := <-errChan:
		return err
	}
}

// Shutdown server
func (s *Server) Shutdown(ctx context.Context) error {
	s.stopChan <- ctx
	return nil
}
