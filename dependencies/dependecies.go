package dependencies

import (
	"github.com/Nivl/go-rest-tools/notifiers/mailer"
	"github.com/Nivl/go-rest-tools/storage/db"
	"github.com/bsphere/le_go"
	"github.com/jmoiron/sqlx"
)

// DB represents an open connection with write access to the database
var DB db.DB

// InitPostgres inits the connection to the database
func InitPostgres(uri string) error {
	con, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return err
	}

	// Unsafe returns a version of DB which will silently succeed to scan when
	// columns in the SQL result have no fields in the destination struct.
	DB = con.Unsafe()
	return nil
}

// Logentries represents an open connection to logentries
var Logentries *le_go.Logger

// InitLogentries inits the connection to logentries
func InitLogentries(token string) {
	le, err := le_go.Connect(token)
	if err != nil {
		panic(err)
	}
	Logentries = le
}

// Sendgrid is a sendgrid email client
var Sendgrid *mailer.Sendgrid

// InitSendgrid creates a mailer that uses Sendgrid
func InitSendgrid(api, from, to, stacktraceUUID string) {
	Sendgrid = mailer.NewSendgrid(api, from, to)
}
