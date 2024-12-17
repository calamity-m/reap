package routes

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/calamity-m/reap/shared/go/middleware"
	"github.com/calamity-m/reap/shared/go/serialize"
)

func NewSowRouter(log *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /echo/", handleEcho(log, "ay"))
	mux.HandleFunc("GET /fail/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "fail", http.StatusBadRequest)
	})

	wrapper := middleware.Wrap(
		middleware.RequestIDMiddleware(log, true),
	)

	return wrapper(mux)
}

func handleEcho(log *slog.Logger, greet string) func(w http.ResponseWriter, r *http.Request) {

	type echo struct {
		Greeting string
		Request  string
		Echo     string
	}

	return func(w http.ResponseWriter, r *http.Request) {

		log.Info("Echo test log, echo handler entered")
		out := &echo{
			Greeting: greet,
			Request:  fmt.Sprintf("%v", r),
			Echo:     "Echhoooo",
		}

		err := serialize.EncodeJSON(w, r, 200, out)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, `{"error": "failed to encode response due to internal server error"}\n`)
			log.Error(fmt.Sprintf("Failed to encode response due to: %v", err.Error()))
		}
	}
}
