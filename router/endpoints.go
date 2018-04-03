package router

import (
	"fmt"
	"net/http"
	"strings"

	reporter "github.com/Nivl/go-reporter"
	"github.com/Nivl/go-rest-tools/network/http/basicauth"
	"github.com/Nivl/go-rest-tools/security/auth"
	"github.com/Nivl/go-rest-tools/types/apperror"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Endpoints represents a list of endpoint
type Endpoints []*Endpoint

// Activate adds the endpoints to the router
func (endpoints Endpoints) Activate(router *mux.Router, deps Dependencies) {
	for _, endpoint := range endpoints {
		router.
			Methods(endpoint.Verb).
			Path(endpoint.Path).
			Handler(Handler(endpoint, deps))
	}
}

// Handler makes it possible to use a RouteHandler where a http.Handler is required
func Handler(e *Endpoint, deps Dependencies) http.Handler {
	HTTPHandler := func(resWriter http.ResponseWriter, req *http.Request) {
		// the following errors will be checked later on. we first init
		// the request, then we will use that request to return (and log) the error
		logger, loggerErr := deps.NewLogger()
		rep, reporterErr := deps.NewReporter()
		request := &HTTPRequest{
			id:       uuid.NewV4().String()[:8],
			http:     req,
			res:      NewResponse(resWriter),
			logger:   logger,
			reporter: rep,
		}
		defer request.handlePanic()

		// We set some response data
		request.res.Header().Set("X-Request-Id", request.id)

		// if a dep failed to be created, we return an error
		if loggerErr != nil {
			request.res.Error(loggerErr, request)
			return
		}
		if reporterErr != nil {
			request.res.Error(reporterErr, request)
			return
		}

		// We setup all the basic tag in the reporter
		request.Reporter().AddTag("Req ID", request.id)
		request.Reporter().AddTag("Endpoint", e.Path)

		// We fetch the user session if a token is provided
		headers, found := req.Header["Authorization"]
		if found {
			request.Reporter().AddTag("Req Auths", strings.Join(headers, ", "))

			userID, sessionID, err := basicauth.ParseAuthHeader(headers, "basic", "")
			if err != nil {
				request.res.Error(apperror.NewBadRequest("Authorization", "invalid format"), request)
			}
			session := &auth.Session{ID: sessionID, UserID: userID}

			if session.ID != "" && session.UserID != "" {
				exists, err := session.Exists(deps.DB())
				if err != nil {
					request.res.Error(err, request)
					return
				}
				if !exists {
					request.res.Error(apperror.NewNotFoundField("Authorization", "session not found"), request)
					return
				}
				request.session = session
				// we get the user and make sure it (still) exists
				request.user, err = auth.GetUserByID(deps.DB(), session.UserID)
				if err != nil {
					if apperror.IsNotFound(err) {
						err = apperror.NewNotFoundField("Authorization", "session not found")
					}
					request.res.Error(err, request)
					return
				}
			}

			request.Reporter().SetUser(&reporter.User{
				ID:       request.user.ID,
				Email:    request.user.Email,
				Username: request.user.Name,
			})
		}

		// Make sure the user has access to the handler
		if allowed, err := e.Guard.HasAccess(request.user); !allowed {
			request.res.Error(err, request)
			return
		}

		// We Parse the request params
		if e.Guard != nil && e.Guard.ParamStruct != nil {
			// Get the list of all http params provided by the client
			sources, err := request.httpParamsBySource()
			if err != nil {
				request.res.Error(err, request)
				return
			}

			request.params, err = e.Guard.ParseParams(sources, request.http)
			if err != nil {
				request.res.Error(err, request)
				return
			}
			request.Reporter().AddTag("Endpoint Params", fmt.Sprintf("%#v", request.params))
		}

		// Execute the actual route handler
		err := e.Handler(request)
		if err != nil {
			request.res.Error(err, request)
		}
	}

	return http.HandlerFunc(HTTPHandler)
}
