package request

import (
	"context"

	logger "github.com/Nivl/go-logger"
	reporter "github.com/Nivl/go-reporter"
	"github.com/Nivl/go-rest-tools/security/auth"
)

// Request represents an http request
//go:generate mockgen -destination mockrequest/request.go -package mockrequest github.com/Nivl/go-rest-tools/request Request
type Request interface {
	String() string

	// Logger returns the logger used by the request
	Logger() logger.Logger

	// Reporter returns the reporter used by the request
	Reporter() reporter.Reporter

	// Signature returns the signature of the request
	// Ex. POST /users
	Signature() string

	// ID returns the ID of the request
	ID() string

	// Response returns the response of the request
	Response() Response

	// Params returns the params needed by the endpoint
	Params() interface{}

	// User returns the user that made the request
	User() *auth.User

	// SetUser sets the user object that made the request
	SetUser(*auth.User)

	// Session returns the session used to make the request
	Session() *auth.Session

	// SetSession sets the session object that was used to make the request
	SetSession(*auth.Session)

	// Context returns the context of the request
	Context() context.Context
}
