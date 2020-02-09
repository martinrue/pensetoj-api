package api

import (
	"fmt"
	"net/http"
	"strconv"
)

type actionRequest struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (r *actionRequest) validate() error {
	validID := false

	truncate := func(s string) string {
		if len(s) > 10 {
			return s[0:10] + "..."
		}

		return s
	}

	for i := 1; i <= 25; i++ {
		if strconv.Itoa(i) == r.ID {
			validID = true
		}
	}

	if !validID {
		return fmt.Errorf("invalid id: %s", truncate(r.ID))
	}

	if r.Type != "like" && r.Type != "listen" {
		return fmt.Errorf("invalid action: %s", truncate(r.Type))
	}

	return nil
}

func (s *Server) handleActions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := s.rid()
		s.addCORSHeaders(w, r)

		var data actionRequest
		if ok := s.readJSON(r, &data); !ok {
			s.writeStatus(w, http.StatusUnprocessableEntity)
			return
		}

		if err := data.validate(); err != nil {
			s.Logger.Print(rid, "invalid request: %s", err)
			s.writeStatus(w, http.StatusUnprocessableEntity)
			return
		}

		s.Logger.Print(rid, "%s: %s", data.Type, data.ID)

		ip, err := s.getRequestIP(r)
		if err != nil {
			s.Logger.Print(rid, "cannot get ip: %s", err)
			s.writeStatus(w, http.StatusInternalServerError)
			return
		}

		if !s.Store.AddAction(data.Type, data.ID, ip) {
			s.Logger.Print(rid, "ignoring, duplicate source")
		}

		s.writeJSON(w, s.Store.GetSummary())
	}
}
