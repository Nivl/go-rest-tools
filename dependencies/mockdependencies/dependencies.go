// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nivl/go-rest-tools/dependencies (interfaces: Dependencies)

// Package mockdependencies is a generated GoMock package.
package mockdependencies

import (
	context "context"
	reflect "reflect"

	go_filestorage "github.com/Nivl/go-rest-tools/vendor/github.com/Nivl/go-filestorage"
	go_logger "github.com/Nivl/go-rest-tools/vendor/github.com/Nivl/go-logger"
	go_mailer "github.com/Nivl/go-rest-tools/vendor/github.com/Nivl/go-mailer"
	go_reporter "github.com/Nivl/go-rest-tools/vendor/github.com/Nivl/go-reporter"
	go_sqldb "github.com/Nivl/go-rest-tools/vendor/github.com/Nivl/go-sqldb"
	gomock "github.com/golang/mock/gomock"
)

// MockDependencies is a mock of Dependencies interface
type MockDependencies struct {
	ctrl     *gomock.Controller
	recorder *MockDependenciesMockRecorder
}

// MockDependenciesMockRecorder is the mock recorder for MockDependencies
type MockDependenciesMockRecorder struct {
	mock *MockDependencies
}

// NewMockDependencies creates a new mock instance
func NewMockDependencies(ctrl *gomock.Controller) *MockDependencies {
	mock := &MockDependencies{ctrl: ctrl}
	mock.recorder = &MockDependenciesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDependencies) EXPECT() *MockDependenciesMockRecorder {
	return m.recorder
}

// DB mocks base method
func (m *MockDependencies) DB() go_sqldb.Connection {
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(go_sqldb.Connection)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockDependenciesMockRecorder) DB() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockDependencies)(nil).DB))
}

// DefaultLogger mocks base method
func (m *MockDependencies) DefaultLogger() (go_logger.Logger, error) {
	ret := m.ctrl.Call(m, "DefaultLogger")
	ret0, _ := ret[0].(go_logger.Logger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultLogger indicates an expected call of DefaultLogger
func (mr *MockDependenciesMockRecorder) DefaultLogger() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultLogger", reflect.TypeOf((*MockDependencies)(nil).DefaultLogger))
}

// Mailer mocks base method
func (m *MockDependencies) Mailer() (go_mailer.Mailer, error) {
	ret := m.ctrl.Call(m, "Mailer")
	ret0, _ := ret[0].(go_mailer.Mailer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mailer indicates an expected call of Mailer
func (mr *MockDependenciesMockRecorder) Mailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mailer", reflect.TypeOf((*MockDependencies)(nil).Mailer))
}

// NewFileStorage mocks base method
func (m *MockDependencies) NewFileStorage(arg0 context.Context) (go_filestorage.FileStorage, error) {
	ret := m.ctrl.Call(m, "NewFileStorage", arg0)
	ret0, _ := ret[0].(go_filestorage.FileStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFileStorage indicates an expected call of NewFileStorage
func (mr *MockDependenciesMockRecorder) NewFileStorage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFileStorage", reflect.TypeOf((*MockDependencies)(nil).NewFileStorage), arg0)
}

// NewLogger mocks base method
func (m *MockDependencies) NewLogger() (go_logger.Logger, error) {
	ret := m.ctrl.Call(m, "NewLogger")
	ret0, _ := ret[0].(go_logger.Logger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewLogger indicates an expected call of NewLogger
func (mr *MockDependenciesMockRecorder) NewLogger() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLogger", reflect.TypeOf((*MockDependencies)(nil).NewLogger))
}

// NewReporter mocks base method
func (m *MockDependencies) NewReporter() (go_reporter.Reporter, error) {
	ret := m.ctrl.Call(m, "NewReporter")
	ret0, _ := ret[0].(go_reporter.Reporter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewReporter indicates an expected call of NewReporter
func (mr *MockDependenciesMockRecorder) NewReporter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReporter", reflect.TypeOf((*MockDependencies)(nil).NewReporter))
}

// SetDB mocks base method
func (m *MockDependencies) SetDB(arg0 go_sqldb.Connection) {
	m.ctrl.Call(m, "SetDB", arg0)
}

// SetDB indicates an expected call of SetDB
func (mr *MockDependenciesMockRecorder) SetDB(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDB", reflect.TypeOf((*MockDependencies)(nil).SetDB), arg0)
}

// SetFileStorageCreator mocks base method
func (m *MockDependencies) SetFileStorageCreator(arg0 go_filestorage.Creator) {
	m.ctrl.Call(m, "SetFileStorageCreator", arg0)
}

// SetFileStorageCreator indicates an expected call of SetFileStorageCreator
func (mr *MockDependenciesMockRecorder) SetFileStorageCreator(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFileStorageCreator", reflect.TypeOf((*MockDependencies)(nil).SetFileStorageCreator), arg0)
}

// SetLoggerCreator mocks base method
func (m *MockDependencies) SetLoggerCreator(arg0 go_logger.Creator) {
	m.ctrl.Call(m, "SetLoggerCreator", arg0)
}

// SetLoggerCreator indicates an expected call of SetLoggerCreator
func (mr *MockDependenciesMockRecorder) SetLoggerCreator(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoggerCreator", reflect.TypeOf((*MockDependencies)(nil).SetLoggerCreator), arg0)
}

// SetMailer mocks base method
func (m *MockDependencies) SetMailer(arg0 go_mailer.Mailer) {
	m.ctrl.Call(m, "SetMailer", arg0)
}

// SetMailer indicates an expected call of SetMailer
func (mr *MockDependenciesMockRecorder) SetMailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMailer", reflect.TypeOf((*MockDependencies)(nil).SetMailer), arg0)
}

// SetReporterCreator mocks base method
func (m *MockDependencies) SetReporterCreator(arg0 go_reporter.Creator) {
	m.ctrl.Call(m, "SetReporterCreator", arg0)
}

// SetReporterCreator indicates an expected call of SetReporterCreator
func (mr *MockDependenciesMockRecorder) SetReporterCreator(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReporterCreator", reflect.TypeOf((*MockDependencies)(nil).SetReporterCreator), arg0)
}
