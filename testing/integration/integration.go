package integration

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"

	db "github.com/Nivl/go-sqldb"
	sqlx "github.com/Nivl/go-sqldb/implementations/sqlxdb"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
	uuid "github.com/satori/go.uuid"
)

type dbManager interface {
	DB() db.Connection
	SetDB(db.Connection)
}

var (
	// dbLock is used to prevent concurent create/update on the temporary table
	dbLock sync.Mutex

	// templateDBCreated is used to check if the template database has already
	// been created
	templateDBCreated = false

	// masterDbName holds the name of the default database
	masterDbName = ""
)

func createTemplateDatabase(con db.Connection, templateName, migrationFolder string) error {
	if templateDBCreated {
		return nil
	}

	// We get the current database name
	if err := con.Get(&masterDbName, "SELECT current_database();"); err != nil {
		return fmt.Errorf("failed getting master database name: %s", err.Error())
	}

	// We drop whatever exists
	stmt := fmt.Sprintf(`DROP DATABASE IF EXISTS "%s";`, templateName)
	if _, err := con.Exec(stmt); err != nil {
		return err
	}

	// We create the template
	stmt = fmt.Sprintf(`CREATE DATABASE "%s";`, templateName)
	if _, err := con.Exec(stmt); err != nil {
		return err
	}

	// We now need to connect to this database to create all the table. We
	// get the DSN of the current connection and swap the table name by
	// the new one
	masterDBString := fmt.Sprintf("dbname=%s", masterDbName)
	tplDBString := fmt.Sprintf("dbname=%s", templateName)
	tplDBDSN := strings.Replace(con.DSN(), masterDBString, tplDBString, -1)
	tplDB, err := sqlx.New(tplDBDSN)
	if err != nil {
		return fmt.Errorf("could not connect to the tmp table: %s", err.Error())
	}
	defer tplDB.Close()

	// We apply the migration to the newly created database
	if err := goose.Up(tplDB.SQL(), migrationFolder); err != nil {
		return err
	}

	templateDBCreated = true
	return nil
}

// New creates a new database, sets it as the current database and returns
// a Wrapper object used to handle panics and clean up databases
func New(manager dbManager, migrationFolder string) (*Wrapper, error) {
	dbLock.Lock()

	templateName := "test_template"
	if err := createTemplateDatabase(manager.DB(), templateName, migrationFolder); err != nil {
		dbLock.Unlock()
		return nil, err
	}
	dbLock.Unlock()

	it := &Wrapper{
		masterDB:  manager.DB(),
		tmpDBName: strings.Replace(uuid.NewV4().String(), "-", "", -1),
	}

	// We create a new Database to avoid races between tests
	stmt := fmt.Sprintf(`CREATE DATABASE "%s" TEMPLATE "%s";`, it.tmpDBName, templateName)
	if _, err := it.masterDB.Exec(stmt); err != nil {
		return nil, errors.Wrap(err, "failed creating tmp database")
	}

	// We set this new database as the new current one, to do that we
	// get the DSN of the current connection and swap the table name by
	// the new one
	masterDBString := fmt.Sprintf("dbname=%s", masterDbName)
	tmpDBString := fmt.Sprintf("dbname=%s", it.tmpDBName)
	tmpDBDSN := strings.Replace(it.masterDB.DSN(), masterDBString, tmpDBString, -1)
	tmpDB, err := sqlx.New(tmpDBDSN)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to the tmp table")
	}
	manager.SetDB(tmpDB)
	it.newDb = manager.DB()

	return it, nil
}

// Wrapper is an helper to simplify the encapsulation of integration
// tests
type Wrapper struct {
	newDb db.Connection
	// masterDB contains a DB connection to the default database.
	// this is needed because postgres won't allow us to drop the
	// current database, so we use the default one to drop the
	// temporary one
	masterDB db.Connection
	// for the new table name we use a uuid without "-"
	tmpDBName string
}

// Close cleans up the tests by deleting the database
func (it *Wrapper) Close() error {
	if err := it.newDb.Close(); err != nil {
		return err
	}
	stmt := fmt.Sprintf(`DROP DATABASE "%s";`, it.tmpDBName)
	if _, err := it.masterDB.Exec(stmt); err != nil {
		return err
	}
	return it.masterDB.Close()
}

// RecoverPanic prevents a panic from not calling the defer in the other goroutines
func (it *Wrapper) RecoverPanic(t *testing.T) {
	if rec := recover(); rec != nil {
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, false)
		t.Fatalf("%v\n%s", rec, string(buf[0:stackSize]))
	}
}

// CloseOnPanic cleans up the tests and re-panic
func (it *Wrapper) CloseOnPanic() {
	if rec := recover(); rec != nil {
		it.Close()
		panic(fmt.Sprintf("%v", rec))
	}
}
