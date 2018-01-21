package router

import (
	filestorage "github.com/Nivl/go-filestorage"
	"github.com/Nivl/go-hasher"
	mailer "github.com/Nivl/go-mailer"
	db "github.com/Nivl/go-sqldb"
)

// Dependencies represents all the dependencies of the API
type Dependencies struct {
	DB      db.Connection
	Mailer  mailer.Mailer
	Storage filestorage.FileStorage
	Hasher  hasher.Hasher
}
