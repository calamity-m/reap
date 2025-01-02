package server

import (
	"log/slog"
	"testing"

	"github.com/calamity-m/reap/sow/internal/config"
)

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
