package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) readJSON(r *http.Request, data interface{}) bool {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&data) == nil
}

func (s *Server) writeJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}
