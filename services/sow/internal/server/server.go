package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/calamity-m/reap/services/sow/internal/routes"
	"github.com/calamity-m/reap/services/sow/internal/service"
)

type SowServer struct {
	srv           http.Server
	log           *slog.Logger
	shutdownGrace time.Duration
}

func NewSowServer(log *slog.Logger, address string) (*SowServer, error) {

	frs, err := service.NewFoodRecorderService()

	if err != nil {
		return nil, err
	}

	srv := &SowServer{
		srv: http.Server{
			Addr:    address,
			Handler: routes.NewSowRouter(log, frs),
		},
		log:           log,
		shutdownGrace: 1,
	}

	return srv, nil
}

func (s *SowServer) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *SowServer) Shutdown() error {

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownGrace*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Error(fmt.Sprintf("Failed to shutdown due to: %v", err))
		return fmt.Errorf("failed to shutdown gracefully: %v", err)
	}

	s.log.Info("successfully completed shutdown")
	return nil
}
