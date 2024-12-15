package routes

import (
	"fmt"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello %v\n", r)
}
