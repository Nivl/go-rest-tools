package router

import (
	"github.com/Nivl/go-rest-tools/dependencies"
	"github.com/Nivl/go-rest-tools/notifiers/mailer"
	"github.com/Nivl/go-rest-tools/storage/db"
)

// Dependencies represents all the dependencies of the API
type Dependencies struct {
	DB     db.DB
	Mailer mailer.Mailer
}

func NewDefaultDependencies() *Dependencies {
	deps := &Dependencies{
		DB:     dependencies.DB,
		Mailer: &mailer.Noop{},
	}

	if dependencies.Sendgrid != nil {
		deps.Mailer = dependencies.Sendgrid
	}

	return deps
}
