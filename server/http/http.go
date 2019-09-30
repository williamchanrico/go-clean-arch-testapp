package http

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/williamchanrico/xtest/listener"
	"github.com/williamchanrico/xtest/xtest"
)

// Server struct
type Server struct {
	address string
	xtest   *xtest.Service

	server *http.Server
}

// New HTTP server
func New(address string, xtestSvc *xtest.Service) *Server {
	return &Server{
		address: address,
		xtest:   xtestSvc,
	}
}

// Run HTTP server
func (s *Server) Run() error {
	l, err := listener.Listen(s.address)
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", s.Ping)
	r.Get("/redis", s.Redis)
	r.Get("/postgres", s.Postgres)
	r.Get("/nsq", s.NSQ)

	httpServer := &http.Server{
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	s.server = httpServer

	return httpServer.Serve(l)
}

// Shutdown HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	if s.server == nil {
		return nil
	}

	return s.server.Shutdown(ctx)
}

// Ping pong
func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong!\n"))
}
