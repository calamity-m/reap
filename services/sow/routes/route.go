package routes

import (
	"fmt"
	"net/http"
)

func NewSowRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /echo/", handleEcho("ay"))

	return mux
}

func handleEcho(greet string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echo echo %s ecchooo... request: %v", greet, r)
	}
}
