package router

import (
	logger "github.com/Nivl/go-logger"
	reporter "github.com/Nivl/go-reporter"
	db "github.com/Nivl/go-sqldb"
)

// Dependencies represents all the dependencies needed by the router
type Dependencies interface {
	// NewLogger creates a new logger using the provided logger creator
	NewLogger() (logger.Logger, error)

	// NewReporter creates a new reporter using the provided reporter Creator
	NewReporter() (reporter.Reporter, error)

	// DB returns the current SQL connection
	DB() db.Connection
}
