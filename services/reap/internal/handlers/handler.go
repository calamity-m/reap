package handlers

import (
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/pkg/middleware"
	"github.com/calamity-m/reap/proto/sow/v1"
)

const (
	SowGetPath         = "GET /food/{userid}/{uuid}"
	SowUpdatePath      = "PUT /food/{userid}/{uuid}"
	SowDeletePath      = "DELETE /food/{userid}/{uuid}"
	SowGetFilteredPath = "GET /food/{userid}"
	SowCreatePath      = "POST /food/{userid}"
)

func NewReaperRouter(log *slog.Logger, sowClient sow.FoodRecordingServiceClient) http.Handler {
	mux := http.NewServeMux()

	// Provision handling of sow related routes
	mux.HandleFunc(SowGetPath, handleSowGet(log, sowClient))
	mux.HandleFunc(SowGetFilteredPath, handleSowGetMany(log, sowClient))
	mux.HandleFunc(SowCreatePath, handleSowCreate(log, sowClient))
	mux.HandleFunc(SowUpdatePath, handleSowUpdate(log, sowClient))
	mux.HandleFunc(SowDeletePath, handleSowDelete(log, sowClient))

	wrapper := middleware.Wrap(
		middleware.RequestIDMiddleware(log, true),
		middleware.LoggingMiddleware(log),
	)

	return wrapper(mux)
}
