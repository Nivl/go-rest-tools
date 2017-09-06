// Code generated by mockery v1.0.0
package mockdependencies

import context "context"
import db "github.com/Nivl/go-rest-tools/storage/db"

import filestorage "github.com/Nivl/go-rest-tools/storage/filestorage"
import logger "github.com/Nivl/go-rest-tools/logger"
import mailer "github.com/Nivl/go-rest-tools/notifiers/mailer"
import mock "github.com/stretchr/testify/mock"

// Dependencies is an autogenerated mock type for the Dependencies type
type Dependencies struct {
	mock.Mock
}

// DB provides a mock function with given fields:
func (_m *Dependencies) DB() db.Connection {
	ret := _m.Called()

	var r0 db.Connection
	if rf, ok := ret.Get(0).(func() db.Connection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Connection)
		}
	}

	return r0
}

// FileStorage provides a mock function with given fields: ctx
func (_m *Dependencies) FileStorage(ctx context.Context) (filestorage.FileStorage, error) {
	ret := _m.Called(ctx)

	var r0 filestorage.FileStorage
	if rf, ok := ret.Get(0).(func(context.Context) filestorage.FileStorage); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filestorage.FileStorage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logger provides a mock function with given fields:
func (_m *Dependencies) Logger() logger.Logger {
	ret := _m.Called()

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func() logger.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// Mailer provides a mock function with given fields:
func (_m *Dependencies) Mailer() mailer.Mailer {
	ret := _m.Called()

	var r0 mailer.Mailer
	if rf, ok := ret.Get(0).(func() mailer.Mailer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mailer.Mailer)
		}
	}

	return r0
}

// SetCloudinary provides a mock function with given fields: apiKey, secret, bucket
func (_m *Dependencies) SetCloudinary(apiKey string, secret string, bucket string) error {
	ret := _m.Called(apiKey, secret, bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(apiKey, secret, bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDB provides a mock function with given fields: uri
func (_m *Dependencies) SetDB(uri string) error {
	ret := _m.Called(uri)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uri)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGCP provides a mock function with given fields: apiKey, projectName, bucket
func (_m *Dependencies) SetGCP(apiKey string, projectName string, bucket string) error {
	ret := _m.Called(apiKey, projectName, bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(apiKey, projectName, bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogentries provides a mock function with given fields: token
func (_m *Dependencies) SetLogentries(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSendgrid provides a mock function with given fields: apiKey, from, to, stacktraceUUID
func (_m *Dependencies) SetSendgrid(apiKey string, from string, to string, stacktraceUUID string) error {
	ret := _m.Called(apiKey, from, to, stacktraceUUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(apiKey, from, to, stacktraceUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
