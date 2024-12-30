package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/reap/pkg/logging"
	"github.com/calamity-m/reap/services/sow/internal/server"
)

func main() {

	logger := slog.New(logging.NewCustomizedHandler(os.Stderr, &logging.CustomHandlerCfg{
		Structed:        false,
		RecordRequestId: true,
		Level:           slog.LevelDebug,
	}))

	sowServer, err := server.NewSowServer("localhost:9098", logger)

	if err != nil {
		logger.Error("failed to create sow rpc server")
		os.Exit(1)
	}

	if err := sowServer.Run(); err != nil {
		logger.Error(fmt.Sprintf("encountered error running sow server: %v", err))
		os.Exit(1)
	}

	logger.Info("quitting")
}
