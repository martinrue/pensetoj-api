package api

import (
	"net/http"
)

func (s *Server) handle404() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Print(s.rid(), "[404] %s", r.URL.Path)
		s.writeStatus(w, http.StatusNotFound)
	}
}
