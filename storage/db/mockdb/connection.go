package mockdb

import db "github.com/Nivl/go-rest-tools/storage/db"
import mock "github.com/stretchr/testify/mock"
import sql "database/sql"

// Connection is a mock type for the Connection type
type Connection struct {
	Queryable
	mock.Mock
}

// Beginx provides a mock function with given fields:
func (_m *Connection) Beginx() (db.Tx, error) {
	ret := _m.Called()

	var r0 db.Tx
	if rf, ok := ret.Get(0).(func() db.Tx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Connection) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DSN provides a mock function with given fields:
func (_m *Connection) DSN() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Exec provides a mock function with given fields: query, arg
func (_m *Connection) Exec(query string, arg ...interface{}) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, arg...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...interface{}) int64); ok {
		r0 = rf(query, arg...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(query, arg...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: dest, query, args
func (_m *Connection) Get(dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamedExec provides a mock function with given fields: query, arg
func (_m *Connection) NamedExec(query string, arg interface{}) (int64, error) {
	ret := _m.Called(query, arg)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, interface{}) int64); ok {
		r0 = rf(query, arg)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamedGet provides a mock function with given fields: dest, query, args
func (_m *Connection) NamedGet(dest interface{}, query string, args interface{}) error {
	ret := _m.Called(dest, query, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, interface{}) error); ok {
		r0 = rf(dest, query, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamedSelect provides a mock function with given fields: dest, query, args
func (_m *Connection) NamedSelect(dest interface{}, query string, args interface{}) error {
	ret := _m.Called(dest, query, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, interface{}) error); ok {
		r0 = rf(dest, query, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SQL provides a mock function with given fields:
func (_m *Connection) SQL() *sql.DB {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// Select provides a mock function with given fields: dest, query, args
func (_m *Connection) Select(dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
