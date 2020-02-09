package api

import (
	"net/http"
)

// Commit is the current Git SHA, injected at build-time.
var Commit = ""

func (s *Server) handleHealthz() http.HandlerFunc {
	type response struct {
		Commit   string `json:"commit"`
		Database string `json:"database"`
		Passed   bool   `json:"passed"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		rid := s.rid()
		s.addCORSHeaders(w, r)

		response := &response{
			Commit: Commit,
			Passed: true,
		}

		size, err := s.Store.Size()
		if err != nil {
			s.Logger.Print(rid, "health: failed to get db size: %v", err)
			response.Passed = false
		}

		response.Database = size

		s.writeJSON(w, response)
	}
}
