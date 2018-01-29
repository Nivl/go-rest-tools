package request

import "net/http"

// Response represents the response of a request
//go:generate mockgen -destination mockrequest/response.go -package mockrequest github.com/Nivl/go-rest-tools/request Response
type Response interface {
	// Header returns the header map that will be attached to the response
	Header() http.Header

	// NoContent sends a response with no content
	NoContent()

	// Created sends a response with a newly created entity attached
	Created(obj interface{}) error

	// Created sends response with a JSON object attached
	Ok(obj interface{}) error
}
