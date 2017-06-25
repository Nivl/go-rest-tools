package router

// RouteHandler is the function signature we nee
type RouteHandler func(*Request, *Dependencies) error

// Endpoint represents an HTTP endpoint
type Endpoint struct {
	Verb string

	// Path is the path for the current component
	Path string

	// Handler is the handler to call
	Handler RouteHandler

	Guard *Guard
}
