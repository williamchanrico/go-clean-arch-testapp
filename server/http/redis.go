package http

import (
	"fmt"
	"net/http"
)

// Redis test redis connection
func (s *Server) Redis(w http.ResponseWriter, r *http.Request) {
	redisOpts := s.redis.Options()
	_, err := s.redis.Ping().Result()

	pingResult := ""
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		pingResult = "Redis Ping failed!"
	} else {
		w.WriteHeader(http.StatusOK)
		pingResult = "Redis Ping success!"
	}

	resp := fmt.Sprintf("%v [%v://%v]\n", pingResult, redisOpts.Network, redisOpts.Addr)
	w.Write([]byte(resp))
}
