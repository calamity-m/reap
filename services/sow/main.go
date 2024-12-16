package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func run(srv *SowServer) error {
	exit := make(chan error)

	go func() {
		// At the end of our function. If no errors were otherwise
		// pushed to this channel, it notifies as a successful shutdown
		defer close(exit)

		sig := make(chan os.Signal, 2)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive a Interrupt or Kill
		<-sig

		fmt.Println("Exiting server")
		if err := srv.Shutdown(); err != nil {
			// Pass error through for return
			exit <- err
			return
		}

	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start/close sow server due to: %v", err)
	}

	fmt.Println("Exiting server!!!")
	return <-exit
}

func main() {

	sow := NewSowServer(*slog.Default(), "localhost:8099")

	err := run(sow)

	if err != nil {
		log.Fatalf("Encountered error running sow server: %v", err)
		os.Exit(1)
	}
}
