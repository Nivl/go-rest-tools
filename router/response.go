package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Nivl/go-rest-tools/network/http/httperr"
)

// HTTPResponse represents an http response
type HTTPResponse interface {
	// Header returns the header map that will be sent by WriteHeader
	Header() http.Header

	// NoContent sends a http.StatusNoContent response
	NoContent()

	// Created sends a http.StatusCreated response with a JSON object attached
	Created(obj interface{}) error
	Ok(obj interface{}) error
}

// Response is a basic implementation of the HTTPResponse that uses a ResponseWriter
type Response struct {
	header http.Header
	writer http.ResponseWriter
	deps   *Dependencies
}

// NewResponse creates a new response
func NewResponse(writer http.ResponseWriter, deps *Dependencies) *Response {
	return &Response{
		writer: writer,
		deps:   deps,
	}
}

// Header sends a http.StatusNoContent response
func (res *Response) Header() http.Header {
	return res.header
}

// NoContent sends a http.StatusNoContent response
func (res *Response) NoContent() {
	res.writer.WriteHeader(http.StatusNoContent)
}

// Created sends a http.StatusCreated response with a JSON object attached
func (res *Response) Created(obj interface{}) error {
	return res.renderJSON(http.StatusCreated, obj)
}

// Ok sends a http.StatusOK response with a JSON object attached
func (res *Response) Ok(obj interface{}) error {
	return res.renderJSON(http.StatusOK, obj)
}

// renderJSON attaches a json object to the response
func (res *Response) renderJSON(code int, obj interface{}) error {
	res.setJSON(code)

	if obj != nil {
		return json.NewEncoder(res.writer).Encode(obj)
	}
	return nil
}

// Error sends an error to the client
// If the error is an instance of HTTPError, the returned code will
// match HTTPError.Code(). It returns a 500 if no code has been set.
func (res *Response) Error(e error, req HTTPRequest) {
	err, casted := e.(*httperr.HTTPError)
	if !casted {
		err = httperr.NewServerError(e.Error()).(*httperr.HTTPError)
	}

	switch err.Code() {
	case http.StatusInternalServerError:
		res.errorJSON(`{"error":"Something went wrong"}`, http.StatusInternalServerError)
	default:
		// Some errors do not need a body
		if err.Error() == "" {
			res.writer.WriteHeader(err.Code())
		} else {
			res.errorJSON(fmt.Sprintf(`{"error":"%s"}`, err.Error()), err.Code())
		}
	}

	req.Logger().Errorf(`code: "%d", message: "%s", %s`, err.Code(), err.Error(), req)

	// We send an email for all server error
	if err.Code() == http.StatusInternalServerError {
		sendEmail := func(stacktrace []byte) {
			err := res.deps.Mailer.SendStackTrace(stacktrace, req.Signature(), err.Error(), req.ID())
			if err != nil {
				req.Logger().Error(err.Error())
			}
		}

		go sendEmail(debug.Stack())
	}
}

// errorJSON set the request content to the specified error message and HTTP code.
// The error message should be valid json.
func (res *Response) errorJSON(err string, code int) {
	res.setJSON(code)
	fmt.Fprintln(res.writer, err)
}

// setJSON set the response to JSON and with the specify HTTP code.
func (res *Response) setJSON(code int) {
	res.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.writer.Header().Set("X-Content-Type-Options", "nosniff")
	res.writer.WriteHeader(code)
}
