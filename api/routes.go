package api

import (
	"net/http"
)

type route struct {
	method  string
	pattern string
	handler http.HandlerFunc
}

func (s *Server) routes() {
	routes := []route{
		{"OPTIONS", "/...", s.handleCORS()},
		{"GET", "/healthz", s.handleHealthz()},
		{"POST", "/actions", s.handleActions()},
		{"GET", "/summary", s.handleSummary()},
	}

	for _, route := range routes {
		s.Router.HandleFunc(route.method, route.pattern, route.handler)
	}

	s.Router.NotFound = s.handle404()
}
