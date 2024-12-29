package errs

import "errors"

var (
	ErrTimeout        = errors.New("timeout")
	ErrInvalidRequest = errors.New("invalid request")
	ErrBadRequest     = errors.New("bad request")
	ErrNotFound       = errors.New("not found")
	ErrBadId          = errors.New("id not formed correctly")
	ErrBadUserId      = errors.New("user id not formed correctly")
	ErrInternal       = errors.New("internal server error")
)
