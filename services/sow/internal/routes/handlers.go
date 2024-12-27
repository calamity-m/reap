package routes

import (
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/pkg/rest"
	"github.com/calamity-m/reap/services/sow/internal/service"
	"github.com/google/uuid"
)

const (
	GetPath         = "GET /{userid}/food/{uuid}"
	UpdatePath      = "PUT /{userid}/food/{uuid}"
	DeletePath      = "DELETE /{userid}/food/{uuid}"
	GetFilteredPath = "GET /{userid}/food/"
	CreatePath      = "POST /{userid}/food/"
)

func handleGet(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Grab the user id from the path as our service requires it for querying
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Grab the record id from the path as our service requires it for querying
		uuid, err := uuid.Parse(r.PathValue("uuid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to parse uuid", slog.String("id", uuid.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Verify if we have some entity matching the provided user id/food record id combo
		record, err := frs.Get(userId, uuid)
		if err != nil {
			// TODO: err type checking here might be helpful
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to get food record", slog.String("id", uuid.String()))
			rest.EncodeMessage(w, http.StatusNotFound, "no record found")
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

		// Happy get
		log.LogAttrs(r.Context(), slog.LevelDebug, "found record", slog.Any("record", record))
		rest.EncodeJSON(w, 200, record)
	}
}

func handleGetFiltered(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Get the user id to force onto the filter
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing user id", slog.String("userid", userId.String()), slog.String("error", err.Error()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Construct the filter from provided request body
		filter, err := rest.DecodeJSON[service.FoodRecord](r)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing body", slog.String("error", err.Error()))
			rest.EncodeMessage(w, http.StatusBadRequest, "filter could not be constructed from passed body")
			return
		}

		// Verify if we have any entities matching the filter
		filter.UserId = userId // set the user id as we want to ignore it in parsing
		records, err := frs.GetFiltered(filter)
		if err != nil {
			log.LogAttrs(
				r.Context(),
				slog.LevelError,
				"error occured while finding records",
				slog.String("error", err.Error()),
				slog.String("userid", userId.String()),
				slog.Any("filter", filter),
			)
			rest.EncodeMessage(w, http.StatusInternalServerError, "internal server error occured")
			return
		}

		// Provide a nice log message in the case we found nothing
		if len(records) == 0 {
			log.LogAttrs(
				r.Context(),
				slog.LevelInfo,
				"no records found",
				slog.String("userid", userId.String()),
				slog.Any("filter", filter),
			)
			rest.EncodeMessage(w, http.StatusNotFound, "no records found")
			return
		}

		// Happy gets
		log.LogAttrs(r.Context(), slog.LevelDebug, "found records", slog.Any("records", records))
		rest.EncodeJSON(w, 200, records)
	}
}

func handleCreate(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Similar to the filter, grab the user id to force onto the created record
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Construct the wanted record from the request body
		record, err := rest.DecodeJSON[service.FoodRecord](r)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing body", slog.String("error", err.Error()))
			rest.EncodeMessage(w, http.StatusBadRequest, "body for record creation could not be parsed")
			return
		}

		// Verify if we can create the record
		record.UserId = userId // set the user id as we want to ignore it in parsing
		created, err := frs.Create(record)
		if err != nil {
			// TODO: handle err type. Maybe the user id not found - etc.
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while creating record", slog.String("error", err.Error()))
			rest.EncodeMessage(w, http.StatusInternalServerError, "failed to create food record")
		}

		// Happy creation
		log.LogAttrs(r.Context(), slog.LevelDebug, "created record", slog.Any("record", record))
		rest.EncodeJSON(w, 200, created)
	}
}

func handleUpdate(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Grab the user id from the path so that like the create, we can forcefully set it
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Grab the uuid of the food record that we want to update
		uuid, err := uuid.Parse(r.PathValue("uuid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to parse uuid", slog.String("id", uuid.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Construct the wanted record
		record, err := rest.DecodeJSON[service.FoodRecord](r)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing body", slog.String("error", err.Error()))
			rest.EncodeMessage(w, http.StatusBadRequest, "body for record creation could not be parsed")
			return
		}

		// Verify if we were able to update the record
		record.UserId = userId // set the user id as we want to ignore it in parsing
		err = frs.Update(uuid, record)
		if err != nil {
			// TODO: handle err type. Maybe the record id wasn't found, user id not found - etc.
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while updating record", slog.String("error", err.Error()))
			rest.EncodeMessage(w, http.StatusInternalServerError, "failed to create food record")
		}

		// Happy updateso
		log.LogAttrs(r.Context(), slog.LevelDebug, "created record", slog.Any("record", record))
		rest.EncodeJSON(w, 200, struct{ message string }{message: "update successful"})

	}
}

func handleDelete(log *slog.Logger, frs *service.FoodRecordService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Grab the userId as our service requires it
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Grab the food record id we want to delete
		uuid, err := uuid.Parse(r.PathValue("uuid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to parse uuid", slog.String("id", uuid.String()))
			rest.EncodeMessage(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Verify if we were able to delete the record
		err = frs.Delete(userId, uuid)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while deleting record", slog.String("error", err.Error()))
			rest.EncodeMessage(w, http.StatusInternalServerError, "failed to delete food record")
		}

		// Happy deleto
		rest.EncodeJSON(w, 200, struct{ message string }{message: "delete successful"})
	}
}
