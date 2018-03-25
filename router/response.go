package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nivl/go-rest-tools/request"
	"github.com/Nivl/go-rest-tools/types/apperror"
)

// ResponseError represents the data sent the client when an error occurs
type ResponseError struct {
	Error string `json:"error,omitempty"`
	Field string `json:"field,omitempty"`
}

var _ request.Response = (*HTTPResponse)(nil)

// HTTPResponse is a basic implementation of the HTTPResponse that uses a ResponseWriter
type HTTPResponse struct {
	writer http.ResponseWriter
}

// NewResponse creates a new response
func NewResponse(writer http.ResponseWriter) *HTTPResponse {
	return &HTTPResponse{
		writer: writer,
	}
}

// Header returns the header object of the response
func (res *HTTPResponse) Header() http.Header {
	return res.writer.Header()
}

// NoContent sends a http.StatusNoContent response
func (res *HTTPResponse) NoContent() {
	res.writer.WriteHeader(http.StatusNoContent)
}

// Created sends a http.StatusCreated response with a JSON object attached
func (res *HTTPResponse) Created(obj interface{}) error {
	return res.renderJSON(http.StatusCreated, obj)
}

// Ok sends a http.StatusOK response with a JSON object attached
func (res *HTTPResponse) Ok(obj interface{}) error {
	return res.renderJSON(http.StatusOK, obj)
}

// renderJSON attaches a json object to the response
func (res *HTTPResponse) renderJSON(code int, obj interface{}) error {
	res.setJSON(code)

	if obj != nil {
		return json.NewEncoder(res.writer).Encode(obj)
	}
	return nil
}

// Error sends an error to the client
// If the error is an instance of HTTPError, the returned code will
// match HTTPError.HTTPStatus(). It returns a 500 if no code has been set.
func (res *HTTPResponse) Error(e error, req request.Request) {
	err := apperror.Convert(e)
	res.errorJSON(err)

	// if the error has a field attached we log it
	field := ""
	if err.Field() != "" {
		field = fmt.Sprintf(`, field: "%s"`, err.Field())
	}
	if req.Logger() != nil {
		req.Logger().Errorf(`code: "%d", httpcode: "%d"%s, message: "%s", %s`, err.StatusCode(), apperror.HTTPStatusCode(err.StatusCode()), field, err.Error(), req)
	}

	// We send a report for all server errors
	if apperror.IsInternalServerError(err) {
		if req.Reporter() != nil {
			req.Reporter().ReportError(err)
		}
	}
}

// errorJSON set the request content to the specified error message and HTTP code.
// The error message should be valid json.
func (res *HTTPResponse) errorJSON(err apperror.Error) {
	httpStatusCode := apperror.HTTPStatusCode(err.StatusCode())
	if err.Error() == "" {
		res.writer.WriteHeader(httpStatusCode)
		return
	}
	resError := &ResponseError{
		Error: err.Error(),
		Field: err.Field(),
	}

	if apperror.IsInternalServerError(err) {
		resError.Error = "Something went wrong"
		resError.Field = ""
	}
	res.renderJSON(httpStatusCode, resError)
}

// setJSON set the response to JSON and with the specify HTTP code.
func (res *HTTPResponse) setJSON(code int) {
	res.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.writer.Header().Set("X-Content-Type-Options", "nosniff")
	res.writer.WriteHeader(code)
}
