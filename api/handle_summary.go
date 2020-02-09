package api

import (
	"net/http"
)

func (s *Server) handleSummary() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.addCORSHeaders(w, r)
		s.writeJSON(w, s.Store.GetSummary())
	}
}
