package errors

import "errors"

type Errorer interface {
	Error() error
}

var ErrBadRosterFormat = errors.New("roster format is not valid")
var ErrMissingId = errors.New("no roster id found")
var ErrBadIdFormat = errors.New("id is not a uuid")
var ErrIdNotFound = errors.New("roster not found")
var ErrCannotCreateEntry = errors.New("unable to save roster")
var ErrFromDatabase = errors.New("error from database")

type EndpointError struct {
	err    error
	status int
}

func NewEndpointError(err error, status int) EndpointError {
	return EndpointError{err, status}
}

func (e EndpointError) Error() string {
	return e.err.Error()
}

func (e EndpointError) StatusCode() int {
	return e.status
}
