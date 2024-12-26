package routes

import (
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/pkg/middleware"
	"github.com/calamity-m/reap/services/sow/internal/service"
)

func NewSowRouter(log *slog.Logger, frs *service.FoodRecordService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(GetIdPath, handleGetId(log, frs))
	mux.HandleFunc(CreatePath, handleCreate(log, frs))
	mux.HandleFunc(UpdatePath, handleUpdate(log, frs))
	mux.HandleFunc(DeletePath, handleDelete(log, frs))

	mux.HandleFunc("GET /fail/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "fail", http.StatusBadRequest)
	})

	wrapper := middleware.Wrap(
		middleware.RequestIDMiddleware(log, true),
		middleware.LoggingMiddleware(log),
	)

	return wrapper(mux)
}
