package jsonServer

import (
	"net/http"
)

func (s *server) Getjson() http.HandlerFunc {
	type request struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		req.Name = "example"
		req.Message = "example 2"

		s.respond(w, r, http.StatusOK, req)
	})
}
