package rest

// Validator is an object that can be validated by some abritrary rule/s
type Validator interface {

	// Validate checks the object and returns any of the found
	// problems in a map, consisting of the offending object's
	// field as the key, and the error
	Validate() (errs map[string]error)
}
