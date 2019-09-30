package http

import (
	"net/http"
)

// NSQ test NSQ connection
func (s *Server) NSQ(w http.ResponseWriter, r *http.Request) {
	var err error
	testResult := ""

	addr := r.URL.Query().Get("addr")
	if addr != "" {
		testResult, err = s.xtest.TestNSQNewAddr(addr)
	} else {
		testResult, err = s.xtest.TestNSQDefaultAddr()
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(testResult))
}
