package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/calamity-m/reap/harvest/internal/config"
	"github.com/calamity-m/reap/harvest/internal/disc"
	"github.com/calamity-m/reap/pkg/logging"
	"github.com/spf13/cobra"
)

var debug bool
var rootCmd = &cobra.Command{
	Use:   "harvest",
	Short: "harvest discord bot",
	Long:  "harvest discord bot",
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
			AddSource:       cfg.LogAddSource,
		}))

		// Create and start the discord bot
		logger.Info("creating discord bot")
		bot, err := disc.NewDiscordBot(logger, cfg)
		if err != nil {
			logger.Error("failed to create discord bot", slog.Any("err", err))
			os.Exit(1)
		}
		defer bot.Close(context.TODO())

		if err := bot.Run(); err != nil {
			logger.Error("failed to run discord bot", slog.Any("err", err))
		}

		logger.Info("exiting")
	},
}

func Execute() error {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Force debug mode on sow server")

	return rootCmd.Execute()
}
