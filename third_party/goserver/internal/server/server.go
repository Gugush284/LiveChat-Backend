package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	Logger *logrus.Logger
}

func NewServer(config *config) *server {
	s := &server{
		router: mux.NewRouter(),
		Logger: logrus.New(),
	}

	s.configureRouter(config)

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Configuration of router ...
func (s *server) configureRouter(config *config) {
}

func (s *server) configureLogger(config *config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	s.Logger.SetLevel(level)

	return nil
}

func Start(config *config) error {

	srv := NewServer(config)

	if err := srv.configureLogger(config); err != nil {
		return err
	}

	srv.Logger.Info("starting api server")
	srv.Logger.Debug()

	return http.ListenAndServe(config.BindAddr, srv)
}
