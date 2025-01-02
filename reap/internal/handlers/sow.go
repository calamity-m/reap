package handlers

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/pkg/errs"
	"github.com/calamity-m/reap/pkg/rest"
	"github.com/calamity-m/reap/proto/sow/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleSowGet(log *slog.Logger, sowClient sow.FoodRecordingClient) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Grab the user id from the path
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Grab the record id from the path as our service requires it for querying
		uuid, err := uuid.Parse(r.PathValue("uuid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to parse uuid", slog.String("id", uuid.String()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Fetch record from client
		record, err := sowClient.GetRecord(r.Context(), &sow.GetRecordRequest{Id: uuid.String()})
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				log.LogAttrs(r.Context(), slog.LevelError, "failed to get food record", slog.String("id", uuid.String()))
				rest.EncodeError(w, http.StatusNotFound, "no record found")
				return
			}

			// We have some other bad error
			log.LogAttrs(r.Context(), slog.LevelError, "internal server error occured", slog.String("id", uuid.String()))
			rest.EncodeError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// Return found record
		log.LogAttrs(r.Context(), slog.LevelDebug, "found record", slog.Any("record", record))
		if err := rest.EncodeJSON(w, 200, record); err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to write response", slog.Any("record", record))
			rest.EncodeError(w, http.StatusInternalServerError, "failed to write response")
		}
	}
}

func handleSowGetMany(log *slog.Logger, sowClient sow.FoodRecordingClient) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Get the user id to force onto the filter
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing user id", slog.String("userid", userId.String()), slog.String("error", err.Error()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Construct the filter from provided request body
		wanted, err := rest.DecodeJSON[*sow.GetRecordsRequest](r)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing body", slog.String("error", err.Error()))
			rest.EncodeError(w, http.StatusBadRequest, "filter could not be constructed from passed body")
			return
		}

		// Create the stream that will fetch records from our sow client
		stream, err := sowClient.GetRecords(r.Context(), wanted)

		// Sanity check the stream was created successfully
		if err != nil {
			log.ErrorContext(r.Context(), "failed to create stream for sow grpc client", slog.Any("error", err))
			rest.EncodeError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// Finally fetch the responses
		responses := make([]*sow.GetRecordsResponse, 0, 1)
		for {
			response, err := stream.Recv()

			// Finished our stream, so exit out now
			if err == io.EOF {
				break
			}

			if err != nil {

				switch status.Code(err) {
				case codes.InvalidArgument:
					log.ErrorContext(r.Context(), "invalid argument received from calling grpc client", slog.Any("error", err))
					rest.EncodeError(w, http.StatusBadRequest, "invalid request")
					return
				case codes.NotFound:
					log.ErrorContext(r.Context(), "not found received from calling grpc client", slog.Any("error", err))
					rest.EncodeError(w, http.StatusBadRequest, "nothing found")
					return
				default:
					log.ErrorContext(r.Context(), "internal server error occured from grpc client", slog.Any("error", err))
					rest.EncodeError(w, http.StatusInternalServerError, "internal server error")
					return
				}
			}

			responses = append(responses, response)
		}

		// Return the found records
		log.LogAttrs(r.Context(), slog.LevelDebug, "found records", slog.Any("response", responses))
		if err := rest.EncodeJSON(w, 200, responses); err != nil {
			rest.EncodeError(w, http.StatusInternalServerError, "failed to write response")
		}
	}
}

func handleSowCreate(log *slog.Logger, sowClient sow.FoodRecordingClient) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Similar to the filter, grab the user id to force onto the created record
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Construct the filter from provided request body
		wanted, err := rest.DecodeJSON[*sow.CreateRecordRequest](r)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing body", slog.String("error", err.Error()))
			rest.EncodeError(w, http.StatusBadRequest, "filter could not be constructed from passed body")
			return
		}

		// Create record in sow
		created, err := sowClient.CreateRecord(r.Context(), wanted)

		if err != nil {
			log.ErrorContext(r.Context(), "failed to create record", slog.Any("error", err))
		}

		log.InfoContext(r.Context(), "successfully created record", slog.Any("created", created))
	}
}

func handleSowUpdate(log *slog.Logger, sowClient sow.FoodRecordingClient) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Grab the user id from the path so that like the create, we can forcefully set it
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Grab the uuid of the food record that we want to update
		uuid, err := uuid.Parse(r.PathValue("uuid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to parse uuid", slog.String("id", uuid.String()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Construct the filter from provided request body
		wanted, err := rest.DecodeJSON[*sow.UpdateRecordRequest](r)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing body", slog.String("error", err.Error()))
			rest.EncodeError(w, http.StatusBadRequest, "filter could not be constructed from passed body")
			return
		}

		// Create record in sow
		updated, err := sowClient.UpdateRecord(r.Context(), wanted)

		if err != nil {
			log.ErrorContext(r.Context(), "failed to update record", slog.Any("error", err))
		}

		log.InfoContext(r.Context(), "successfully updated record", slog.Any("updated", updated))

	}
}

func handleSowDelete(log *slog.Logger, sowClient sow.FoodRecordingClient) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		// Grab the userId as our service requires it
		userId, err := uuid.Parse(r.PathValue("userid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "failed to parse user uuid", slog.String("userid", userId.String()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Grab the food record id we want to delete
		uuid, err := uuid.Parse(r.PathValue("uuid"))
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "Failed to parse uuid", slog.String("id", uuid.String()))
			rest.EncodeError(w, http.StatusBadRequest, "invalid path")
			return
		}

		// Construct the filter from provided request body
		toDelete, err := rest.DecodeJSON[*sow.DeleteRecordRequest](r)
		if err != nil {
			log.LogAttrs(r.Context(), slog.LevelError, "error occured while parsing body", slog.String("error", err.Error()))
			rest.EncodeError(w, http.StatusBadRequest, "filter could not be constructed from passed body")
			return
		}

		// Create record in sow
		deleted, err := sowClient.DeleteRecord(r.Context(), toDelete)

		if err != nil {
			log.ErrorContext(r.Context(), "failed to delete record", slog.Any("error", err))
		}

		log.InfoContext(r.Context(), "successfully deleted record", slog.Any("deleted", deleted))
	}
}
