package routes

import (
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/pkg/rest"
	"github.com/calamity-m/reap/services/sow/internal/service"
	"github.com/google/uuid"
)

const (
	GetPath         = "GET {userid}/food/{uuid}"
	GetFilteredPath = "GET {userid}/food/"
	CreatePath      = "POST {userid}/food/"
	DeletePath      = "DELETE {userid}/food/"
	UpdatePath      = "PUT {userid}/food/"
)

func handleGet(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		userId, err := uuid.Parse(r.PathValue("userid"))

		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		uuid, err := uuid.Parse(r.PathValue("uuid"))

		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to parse uuid", slog.String("id", uuid.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		record, err := frs.Get(userId, uuid)

		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to get food record", slog.String("id", uuid.String()))
			rest.EncodeMessage(w, http.StatusNotFound, "Could not find record")
			return
		}

		log.LogAttrs(
			r.Context(),
			slog.LevelInfo,
			"Found record",
			slog.Group(
				// Because we omit some values as part of the food record struct
				// we capture this here as a separate group broken out instead of
				// using slog.Any("record", record)
				"record",
				slog.String("uuid", record.Uuid.String()),
				slog.String("userid", record.UserId.String()),
				slog.String("name", record.Name),
				slog.String("description", record.Description),
				slog.Float64("kj", float64(record.KJ)),
				slog.Float64("gram", float64(record.Gram)),
				slog.Float64("ml", float64(record.ML)),
				slog.Time("created", record.Created),
			),
		)

		rest.EncodeJSON(w, 200, record)

	}
}

func handleGetFiltered(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO GET FILTERED")

	}
}

func handleCreate(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// body, err := rest.DecodeJSON[service.FoodRecord](r)

		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO CREATE")

	}
}

func handleUpdate(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO UPDATE")

	}
}

func handleDelete(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.InfoContext(r.Context(), "TODO")

		rest.EncodeMessage(w, 500, "TODO DELETE")

	}
}