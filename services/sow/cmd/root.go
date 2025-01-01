package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/reap/pkg/logging"
	"github.com/calamity-m/reap/services/sow/internal/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sow",
	Short: "sow grpc server",
	Long:  "sow grpc server for food recording within reap",
	Run: func(cmd *cobra.Command, args []string) {

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
	},
}

func Execute() error {
	return rootCmd.Execute()
}
