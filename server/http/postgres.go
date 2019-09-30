package http

import (
	"net/http"
)

// Postgres test postgresql db connection
func (s *Server) Postgres(w http.ResponseWriter, r *http.Request) {
	testResult := ""
	var err error

	dsn := r.URL.Query().Get("dsn")
	if dsn != "" {
		testResult, err = s.xtest.TestPostgresNewDSN(dsn)
	} else {
		testResult, err = s.xtest.TestPostgresDefaultDSN()
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(testResult))
}
