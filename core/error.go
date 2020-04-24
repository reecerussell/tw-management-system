package core

// Error is a custom error type, used to more descriptively define errors.
type Error interface {
	Err() error
	Message() string
	Status() int
}

// NewError returns a new instance of Error.
func NewError(err error) Error {
	return &basicError{
		err: err,
	}
}

// basicError is just a wrapper arround error.
type basicError struct {
	err error
}

// Err returns the underlying error.
func (be *basicError) Err() error {
	return be.err
}

// Message returns the error message.
func (be *basicError) Message() string {
	return be.err.Error()
}

// Status returns a 500 error code.
func (be *basicError) Status() int {
	return 500
}

// NewErrorWithStatus returns a new Error with the given status code and error.
func NewErrorWithStatus(err error, status int) Error {
	return &statusError{
		err:    err,
		status: status,
	}
}

// statusError wraps the error type and holds a status code for it.
type statusError struct {
	err    error
	status int
}

// Err returns the underlying error.
func (se *statusError) Err() error {
	return se.err
}

// Message returns the error message.
func (se *statusError) Message() string {
	return se.err.Error()
}

// Status returns the error's status code.
func (se *statusError) Status() int {
	return se.status
}
