package routes

import (
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/pkg/rest"
)

func handleGet(log *slog.Logger) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO GET")

	}
}

func handleCreate(log *slog.Logger) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO CREATE")

	}
}

func handleUpdate(log *slog.Logger) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO UPDATE")

	}
}

func handleDelete(log *slog.Logger) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO DELETE")

	}
}
