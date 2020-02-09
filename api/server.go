package api

import (
	"net/http"

	"github.com/matryer/way"

	"github.com/martinrue/pensetoj-api/logger"
	"github.com/martinrue/pensetoj-api/store"
)

// Server defines the API server and its dependencies.
type Server struct {
	Logger *logger.Logger
	Router *way.Router
	Store  store.Store
}

// Start brings up the server.
func (s *Server) Start(bind string) error {
	s.routes()
	return http.ListenAndServe(bind, s.Router)
}

func (s *Server) writeStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
