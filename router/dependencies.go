package router

import (
	mailer "github.com/Nivl/go-mailer"
	"github.com/Nivl/go-rest-tools/storage/db"
	filestorage "github.com/Nivl/go-filestorage"
)

// Dependencies represents all the dependencies of the API
type Dependencies struct {
	DB      db.Connection
	Mailer  mailer.Mailer
	Storage filestorage.FileStorage
}
