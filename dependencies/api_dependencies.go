package dependencies

import (
	"context"
	"errors"
	"sync"

	"github.com/Nivl/go-rest-tools/logger"
	"github.com/Nivl/go-rest-tools/notifiers/mailer"
	"github.com/Nivl/go-rest-tools/notifiers/reporter"
	"github.com/Nivl/go-rest-tools/storage/db"
	"github.com/Nivl/go-rest-tools/storage/db/sqlx"
	"github.com/Nivl/go-rest-tools/storage/filestorage"
)

var _ Dependencies = (*APIDependencies)(nil)

// APIDependencies is a structure implementing the Dependencies interface
type APIDependencies struct {
	sync.Mutex

	// logentries *le_go.Logger
	// sendgrid   mailer.Mailer
	// gcp        gcp.GCP
	// cloudinary filestorage.FileStorage

	defaultSQLCon *sqlx.Connection
	defaultLogger logger.Logger
	defaultMailer mailer.Mailer

	reporterCreator    reporter.Creator
	loggerCreator      logger.Creator
	filestorageCreator filestorage.Creator
}

// SetDB creates a connection to a SQL database
func (deps *APIDependencies) SetDB(uri string) error {
	deps.Lock()
	defer deps.Unlock()

	var err error
	deps.defaultSQLCon, err = sqlx.New(uri)
	return err
}

// DB returns the current SQL connection
func (deps *APIDependencies) DB() db.Connection {
	return deps.defaultSQLCon
}

// SetLoggerCreator sets a logger creator used to generate new loggers
func (deps *APIDependencies) SetLoggerCreator(creator logger.Creator) {
	deps.loggerCreator = creator
}

// DefaultLogger returns the default logger to use app wide
func (deps *APIDependencies) DefaultLogger() (logger.Logger, error) {
	deps.Lock()
	defer deps.Unlock()

	if deps.loggerCreator == nil {
		return nil, errors.New("no logger creator has been set")
	}
	if deps.defaultLogger == nil {
		var err error
		deps.defaultLogger, err = deps.loggerCreator.New()
		return deps.defaultLogger, err
	}
	return deps.defaultLogger, nil
}

// NewLogger returns a new logger using the logger Creator
func (deps *APIDependencies) NewLogger() (logger.Logger, error) {
	if deps.loggerCreator == nil {
		return nil, errors.New("no logger creator has been set")
	}
	return deps.loggerCreator.New()
}

// SetMailer sets the mailer to be used to send emails
func (deps *APIDependencies) SetMailer(m mailer.Mailer) {
	deps.defaultMailer = m
}

// Mailer returns the mailer set with SetMailer
func (deps *APIDependencies) Mailer() (mailer.Mailer, error) {
	if deps.defaultMailer == nil {
		return nil, errors.New("no mailer has been set")
	}
	return deps.defaultMailer, nil
}

// SetReporterCreator sets a reporter creator used to generate new reporters
func (deps *APIDependencies) SetReporterCreator(creator reporter.Creator) {
	deps.reporterCreator = creator
}

// NewReporter creates a new reporter using the provided reporter Creator
func (deps *APIDependencies) NewReporter() (reporter.Reporter, error) {
	if deps.reporterCreator == nil {
		return nil, errors.New("no reporter creator has been set")
	}
	return deps.reporterCreator.New()
}

// SetFileStorageCreator returns the default filestorage following this order
func (deps *APIDependencies) SetFileStorageCreator(creator filestorage.Creator) {
	deps.filestorageCreator = creator
}

// NewFileStorage creates a new filestorage using the provided reporter Creator
func (deps *APIDependencies) NewFileStorage(ctx context.Context) (filestorage.FileStorage, error) {
	if deps.reporterCreator == nil {
		return nil, errors.New("no filestorage creator has been set")
	}
	return deps.filestorageCreator.New(ctx)
}
