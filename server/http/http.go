package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	address    string
	redis      *redis.Client
	httpServer *http.Server
}

// New HTTP server
func New(address string, redis *redis.Client) *Server {
	return &Server{
		address: address,
		redis:   redis,
	}
}

// Run HTTP server
func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	r := mux.NewRouter()

	r.HandleFunc("/ping", s.Ping).Methods("GET")

	r.HandleFunc("/redis", s.Redis).Methods("GET")

	httpServer := &http.Server{
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	s.httpServer = httpServer

	return httpServer.Serve(l)
}

// Shutdown HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}

	return s.httpServer.Shutdown(ctx)
}

// Ping pong
func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong!\n"))
}
