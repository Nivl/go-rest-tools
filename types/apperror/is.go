package apperror

// IsNotFound checks if an error is the NotFound type
func IsNotFound(e error) bool {
	err, casted := e.(*AppError)
	if !casted {
		return false
	}
	return err.StatusCode() == NotFound
}

// IsConflict checks if an error is caused by a conflict
func IsConflict(e error) bool {
	err, casted := e.(*AppError)
	if !casted {
		return false
	}
	return err.StatusCode() == AlreadyExists
}

// IsInternalServerError checks if an error is caused by an internal error
func IsInternalServerError(e error) bool {
	err := Convert(e)
	return err.StatusCode() == Internal
}

// IsBadRequest checks if an error is caused by a bad request
func IsBadRequest(e error) bool {
	err, casted := e.(*AppError)
	if !casted {
		return false
	}
	return err.StatusCode() == InvalidArgument
}

// IsForbidden checks if an error is caused by a forbidden access
func IsForbidden(e error) bool {
	err, casted := e.(*AppError)
	if !casted {
		return false
	}
	return err.StatusCode() == PermissionDenied
}

// IsUnauthorized checks if an error is caused by an Unauthorized access
func IsUnauthorized(e error) bool {
	err, casted := e.(*AppError)
	if !casted {
		return false
	}
	return err.StatusCode() == Unauthenticated
}

// IsInvalidParam checks if an error is caused by an invalid param
func IsInvalidParam(e error) bool {
	return IsBadRequest(e)
}
