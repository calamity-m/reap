package server

import (
	"log/slog"
	"testing"

	"github.com/calamity-m/reap/sow/internal/config"
)

/*

func NewSowServer(cfg *config.Config, logger *slog.Logger) (*SowGRPCServer, error) {
	if cfg == nil || logger == nil {
		return nil, fmt.Errorf("nil input not allowed")
	}

	store := persistence.NewMemoryFoodStore()
	foodService, err := service.NewFoodRecorderService(logger, store)

	if err != nil {
		logger.Error("failed to create sow food service and store")
		return nil, errors.New("failed to create server")
	}

	server := &SowGRPCServer{
		log:     logger,
		service: foodService,
		addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	return server, nil
}

*/

func TestNewSowServer(t *testing.T) {
	t.Run("cfg cannot be nil", func(t *testing.T) {
		_, err := NewSowServer(nil, slog.Default())

		if err == nil {
			t.Errorf("expected error on nil config")
		}
	})

	t.Run("logger cannont be nil", func(t *testing.T) {
		_, err := NewSowServer(&config.Config{}, nil)

		if err == nil {
			t.Errorf("expected error on nil logger")
		}
	})

	t.Run("server returned with no nils", func(t *testing.T) {
		server, err := NewSowServer(&config.Config{}, slog.Default())

		if err != nil {
			t.Errorf("expected no error")
		}

		if server == nil {
			t.Errorf("did not expect nil server")
		}
	})

	t.Run("server address is correctly formed", func(t *testing.T) {
		want := "test:1000"
		server, err := NewSowServer(&config.Config{Host: "test", Port: 1000}, slog.Default())

		if err != nil {
			t.Errorf("expected no error")
		}

		if server.addr != want {
			t.Errorf("got %s address but wanted %s", server.addr, want)
		}
	})
}
