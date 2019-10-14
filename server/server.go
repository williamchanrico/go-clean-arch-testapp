package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/williamchanrico/xtest/log"
	xtestGRPC "github.com/williamchanrico/xtest/server/grpc"
	xtestHTTP "github.com/williamchanrico/xtest/server/http"
	"github.com/williamchanrico/xtest/xtest"
)

// Server struct
type Server struct {
	HTTPAddress string
	GRPCAddress string
	Xtest       *xtest.Service

	stopChan chan context.Context
}

// Run xtest servers
func (s *Server) Run(ctx context.Context) error {
	errChan := make(chan error)
	s.stopChan = make(chan context.Context)
	httpServer := xtestHTTP.New(s.HTTPAddress, s.Xtest)
	grpcServer := xtestGRPC.New(s.GRPCAddress)

	// Start HTTP server
	go func(httpServer *xtestHTTP.Server) {
		log.Infof("HTTP server is listening on %v\n", s.HTTPAddress)
		err := httpServer.Run()
		errChan <- err
	}(httpServer)

	// Start GRPC server
	go func(grpcServer *xtestGRPC.Server) {
		log.Infof("GRPC server is listening on %v\n", s.GRPCAddress)
		err := grpcServer.Run()
		errChan <- err
	}(grpcServer)

	go func() {
		signals := make(chan os.Signal, 1)

		// Also traps SIGHUP in case init system sends one
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		select {
		case <-signals:
		case <-ctx.Done():
		}

		// Received an OS signal, stahpp the server!
		log.Infof("Shutting down all servers")
		if err := s.Shutdown(ctx); err != nil {
			log.Error("All servers shutdown: ", err.Error())
		}
	}()

	select {
	case ctx := <-s.stopChan:
		grpcServer.Shutdown(ctx)
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
