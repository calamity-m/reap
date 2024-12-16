package serialize

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeJSON[T any](w http.ResponseWriter, r *http.Request, status int, val T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(val); err != nil {
		return fmt.Errorf("failed encoding json: %v", err)
	}

	return nil
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var val T

	if err := json.NewDecoder(r.Body).Decode(&val); err != nil {
		return val, fmt.Errorf("failed decoding json: %v", err)
	}

	return val, nil
}
