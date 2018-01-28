package apperror

// Convert takes an error an turns it into an AppError
func Convert(e error) *AppError {
	err, casted := e.(*AppError)
	if !casted {
		err = NewServerError(e.Error())
		err.origin = e
	}

	return err
}

// Error represents an error with a code attached.
type Error interface {
	error

	// StatusCode return the HTTP code of the error
	StatusCode() Code

	// Field returns the http/sql/etc. field associated to the error
	Field() string

	// Origin returns the original error if there's one
	Origin() error
}

// AppError represents an error with a status code and an optional field
type AppError struct {
	error
	status Code
	field  string
	origin error
}

// StatusCode returns the HTTP code associated to the error
func (err *AppError) StatusCode() Code {
	if err == nil {
		return Internal
	}
	return err.status
}

// Field returns the HTTP param associated to the error
func (err *AppError) Field() string {
	if err == nil {
		return ""
	}
	return err.field
}

// Origin returns the original error
func (err *AppError) Origin() error {
	if err == nil {
		return nil
	}
	return err.origin
}
