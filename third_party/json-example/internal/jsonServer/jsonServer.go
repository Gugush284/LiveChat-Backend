package jsonServer

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	Logger *logrus.Logger
}

func NewServer(config *jsonConfig) *server {
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
func (s *server) configureRouter(config *jsonConfig) {
	s.router.HandleFunc("/getjson", s.Getjson(config)).Methods("GET")
}

func (s *server) configureLogger(config *jsonConfig) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	s.Logger.SetLevel(level)

	return nil
}

func Start(config *jsonConfig) error {

	srv := NewServer(config)

	if err := srv.configureLogger(config); err != nil {
		return err
	}

	srv.Logger.Info("starting api server")
	srv.Logger.Debug(config.SessionKey)

	return http.ListenAndServe(config.BindAddr, srv)
}
