package jsonServer

import (
	"net/http"

	"github.com/BurntSushi/toml"
)

func (s *server) Getjson(config *jsonConfig) http.HandlerFunc {
	type data struct {
		Name    string `toml:"username"`
		Message string `toml:"message"`
	}

	type request struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		d := &data{}

		_, err := toml.DecodeFile(config.DataPath, d)
		if err != nil {
			s.Logger.Info(err)
		}

		req.Name = d.Name
		req.Message = d.Message

		s.respond(w, r, http.StatusOK, req)
	})
}
