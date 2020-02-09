package api

import (
	"net"
	"net/http"

	"github.com/martinrue/pensetoj-api/token"
)

func (s *Server) rid() string {
	return token.NewShort()
}

func (s *Server) addCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
}

func (s *Server) getRequestIP(r *http.Request) (string, error) {
	address := r.Header.Get("X-Real-Ip")
	if address != "" {
		return address, nil
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	return ip, err
}
