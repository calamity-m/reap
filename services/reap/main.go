package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/calamity-m/reap/pkg/logging"
	"github.com/calamity-m/reap/services/reap/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func run(srv *server.ReaperServer) error {
	if srv == nil {
		return fmt.Errorf("cannot run with a nil server")
	}

	exit := make(chan error)

	go func() {
		// At the end of our function. If no errors were otherwise
		// pushed to this channel, it notifies as a successful shutdown
		defer close(exit)

		sig := make(chan os.Signal, 2)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive a Interrupt or Kill
		<-sig

		if err := srv.Shutdown(); err != nil {
			// Pass error through for return
			exit <- err
			return
		}

	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start/close sow server due to: %w", err)
	}

	return <-exit
}

func main() {

	logger := slog.New(logging.NewCustomizedHandler(os.Stderr, &logging.CustomHandlerCfg{
		Structed:        false,
		RecordRequestId: true,
		Level:           slog.LevelDebug,
	}))

	// Create GRPC clients
	sowClient, sowConn, err := createSowClient("localhost:9099", []grpc.DialOption{
		// For now just use insecure
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		logger.Error(fmt.Sprintf("encountered error creating sow grpc client: %v", err))
		os.Exit(1)
	}
	defer sowConn.Close()

	// Create the reap server itself
	reap, err := server.NewReaperServer(logger, "localhost:8099", sowClient)
	if err != nil {
		logger.Error(fmt.Sprintf("Encountered error creating reap server: %v", err))
		// force the closure of grpc clients
		sowConn.Close()

		// Exit now
		os.Exit(1)
	}

	logger.Info("Initialized server, moving to initiating http listen")
	if err := run(reap); err != nil {
		logger.Error(fmt.Sprintf("Encountered error running reap server: %v", err))
		// force the closure of grpc clients
		sowConn.Close()

		// Exit now
		os.Exit(1)
	}

	logger.Info("Exiting reap server")
}
