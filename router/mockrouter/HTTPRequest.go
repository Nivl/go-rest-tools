// Code generated by mockery v1.0.0
package mocks

import auth "github.com/Nivl/go-rest-tools/security/auth"
import logger "github.com/Nivl/go-rest-tools/logger"
import mock "github.com/stretchr/testify/mock"
import router "github.com/Nivl/go-rest-tools/router"

// HTTPRequest is an autogenerated mock type for the HTTPRequest type
type HTTPRequest struct {
	mock.Mock
}

// ID provides a mock function with given fields:
func (_m *HTTPRequest) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *HTTPRequest) Logger() logger.Logger {
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

// Params provides a mock function with given fields:
func (_m *HTTPRequest) Params() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Response provides a mock function with given fields:
func (_m *HTTPRequest) Response() router.HTTPResponse {
	ret := _m.Called()

	var r0 router.HTTPResponse
	if rf, ok := ret.Get(0).(func() router.HTTPResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(router.HTTPResponse)
		}
	}

	return r0
}

// Signature provides a mock function with given fields:
func (_m *HTTPRequest) Signature() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *HTTPRequest) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// User provides a mock function with given fields:
func (_m *HTTPRequest) User() *auth.User {
	ret := _m.Called()

	var r0 *auth.User
	if rf, ok := ret.Get(0).(func() *auth.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.User)
		}
	}

	return r0
}
