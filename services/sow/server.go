package main

import (
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/services/sow/routes"
)

type SowServer struct {
	srv http.Server
	log slog.Logger
}

func (s *SowServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func NewSowServer(log slog.Logger, address string) *SowServer {
	srv := &SowServer{
		srv: http.Server{
			Addr:    address,
			Handler: routes.NewSowRouter(),
		},
		log: log,
	}

	return srv
}

func (s *SowServer) ListenAndServe() error {
	return s.srv.ListenAndServe()
}
