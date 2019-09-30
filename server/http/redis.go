package http

import (
	"net/http"
)

// Redis test redis connection
func (s *Server) Redis(w http.ResponseWriter, r *http.Request) {
	testResult, err := s.xtest.TestRedis()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(testResult))
}
