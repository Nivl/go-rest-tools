package db

import "github.com/jmoiron/sqlx"

// Get is the same as sqlx.Get() but do not returns an error on empty results
func Get(dest interface{}, query string, args ...interface{}) error {
	err := Writer.Get(dest, query, args...)
	if IsNotFound(err) {
		return nil
	}
	return err
}

// NamedSelect is the same as sqlx.Select() but with named params
func NamedSelect(dest interface{}, query string, args interface{}) error {
	row, err := Writer.NamedQuery(query, args)
	if err != nil {
		return err
	}
	return sqlx.StructScan(row, dest)
}
