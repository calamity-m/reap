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

// Decodes an object into a type T, ensuring it passes validation. If the length
// of the error map is > 0, the decoding is considered invalid.
//
// When decoding fails the error map will have a single entry that corresponds
// to the decoding error.
//
// When the decoding passes, but validation fails, the map will contain one or more
// entries corresponding to the validation errors, with each key mapping to the
// field that failed validation
func DecodeJSONValid[T Validator](r *http.Request) (T, map[string]error) {

	val, err := DecodeJSON[T](r)
	if err != nil {
		return val, map[string]error{"decoding": err}
	}

	errs := val.Validate()

	if len(errs) > 0 {
		return val, errs
	}

	return val, nil
}
