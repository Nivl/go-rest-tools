package apperror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// Code represent an Error code
type Code uint

const (
	// InvalidArgument is returned when a user provided data is invalid
	InvalidArgument Code = 100

	// NotFound indicates a requested entity was not found
	NotFound = 101

	// AlreadyExists indicates an attempt to create an entity failed because
	// it already exists
	AlreadyExists = 102

	// Unauthenticated indicates the request does not have a valid
	// authentication credentials
	Unauthenticated = 103

	// PermissionDenied indicates the requester does not have the right
	// permissions to execute the request
	PermissionDenied = 104

	// Internal indicates something the service is internally broken
	Internal = 1000
)

var statusText = map[Code]string{
	InvalidArgument:  "Bad Request",
	Unauthenticated:  "Unauthorized",
	PermissionDenied: "Forbidden",
	NotFound:         "Not Found",
	AlreadyExists:    "Conflict",
	Internal:         "Internal Error",
}

// StatusText returns
func StatusText(errCode Code) string {
	return statusText[errCode]
}

var httpCodes = map[Code]int{
	InvalidArgument:  http.StatusBadRequest,
	Unauthenticated:  http.StatusUnauthorized,
	PermissionDenied: http.StatusForbidden,
	NotFound:         http.StatusNotFound,
	AlreadyExists:    http.StatusConflict,
	Internal:         http.StatusInternalServerError,
}

// HTTPStatusCode returns the HTTP Code corresponding to the
// provided app error status code
func HTTPStatusCode(code Code) int {
	return httpCodes[code]
}

var grpcCodes = map[Code]codes.Code{
	InvalidArgument:  codes.InvalidArgument,
	Unauthenticated:  codes.Unauthenticated,
	PermissionDenied: codes.PermissionDenied,
	NotFound:         codes.NotFound,
	AlreadyExists:    codes.AlreadyExists,
	Internal:         codes.Internal,
}

// GRPCStatusCode returns the GRPC Code corresponding to the
// provided app error status code
func GRPCStatusCode(code Code) codes.Code {
	return grpcCodes[code]
}
