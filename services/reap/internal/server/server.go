package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/calamity-m/reap/proto/sow/v1"
	"github.com/calamity-m/reap/services/reap/internal/handlers"
)

type ReaperServer struct {
	srv           http.Server
	log           *slog.Logger
	shutdownGrace time.Duration
}

func NewReaperServer(log *slog.Logger, address string, sowClient sow.FoodRecordingServiceClient) (*ReaperServer, error) {
	srv := &ReaperServer{
		srv: http.Server{
			Addr:    address,
			Handler: handlers.NewReaperRouter(log, sowClient),
		},
		log:           log,
		shutdownGrace: 1,
	}

	return srv, nil
}

func (s *ReaperServer) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *ReaperServer) Shutdown() error {

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownGrace*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Error(fmt.Sprintf("Failed to shutdown due to: %v", err))
		return fmt.Errorf("failed to shutdown gracefully: %w", err)
	}

	s.log.Info("successfully completed shutdown")
	return nil
}
