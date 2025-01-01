package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/reap/pkg/logging"
	"github.com/calamity-m/reap/sow/config"
	"github.com/calamity-m/reap/sow/internal/server"
	"github.com/spf13/cobra"
)

var debug bool
var rootCmd = &cobra.Command{
	Use:   "sow",
	Short: "sow grpc server",
	Long:  "sow grpc server for food recording within reap",
	Run: func(cmd *cobra.Command, args []string) {

		// Grab config
		cfg, err := config.NewConfig(debug)
		if err != nil {
			fmt.Printf("Failed to create config: %v\n", err)
			os.Exit(1)
		}

		// Create logger
		logger := slog.New(logging.NewCustomizedHandler(os.Stderr, &logging.CustomHandlerCfg{
			Structed:        cfg.LogStructured,
			RecordRequestId: cfg.LogRequestId,
			Level:           cfg.LogLevel,
		}))

		// Create sow server
		sowServer, err := server.NewSowServer(cfg, logger)
		if err != nil {
			logger.Error("failed to create sow rpc server")
			os.Exit(1)
		}

		logger.Debug("a")
		logger.Warn("b")
		// Block on run until we want to quit
		if err := sowServer.Run(); err != nil {
			logger.Error(fmt.Sprintf("encountered error running sow server: %v", err))
			os.Exit(1)
		}

		logger.Info("quitting")
	},
}

func Execute() error {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Force debug mode on sow server")

	return rootCmd.Execute()
}
