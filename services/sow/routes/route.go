package routes

import (
	"fmt"
	"net/http"
)

func NewSowRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /echo/", echo)

	return mux
}

func echo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello %v\n", r)
}
