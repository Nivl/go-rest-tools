package dependencies

import (
	"context"

	logger "github.com/Nivl/go-logger"
	mailer "github.com/Nivl/go-mailer"
	reporter "github.com/Nivl/go-reporter"
	"github.com/Nivl/go-rest-tools/storage/db"
	filestorage "github.com/Nivl/go-filestorage"
)

// Dependencies represents the dependency of the api
type Dependencies interface {
	// SetDB creates a connection to a SQL database
	SetDB(db db.Connection)

	// DB returns the current SQL connection
	DB() db.Connection

	// SetLoggerCreator sets a logger creator used to generate new loggers
	SetLoggerCreator(logger.Creator)

	// NewLogger creates a new logger using the provided logger creator
	NewLogger() (logger.Logger, error)

	// DefaultLogger return a app-wide logger created using the provided
	// logger creator
	DefaultLogger() (logger.Logger, error)

	// SetMailer sets the mailer to be used to send emails
	SetMailer(mailer.Mailer)

	// Mailer returns the mailer set with SetMailer
	Mailer() (mailer.Mailer, error)

	// SetReporterCreator sets a reporter creator used to generate new reporters
	SetReporterCreator(reporter.Creator)

	// NewReporter creates a new reporter using the provided reporter Creator
	NewReporter() (reporter.Reporter, error)

	// SetFileStorageCreator sets a filestorage creator used to generate new
	// filestorage
	SetFileStorageCreator(filestorage.Creator)

	// NewFileStorage creates a new filestorage using the provided reporter Creator
	NewFileStorage(context.Context) (filestorage.FileStorage, error)
}
