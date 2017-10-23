// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nivl/go-rest-tools/security/hasher (interfaces: Hasher)

// Package mockhasher is a generated GoMock package.
package mockhasher

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHasher is a mock of Hasher interface
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// Hash mocks base method
func (m *MockHasher) Hash(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "Hash", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash
func (mr *MockHasherMockRecorder) Hash(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockHasher)(nil).Hash), arg0)
}

// IsValid mocks base method
func (m *MockHasher) IsValid(arg0, arg1 string) bool {
	ret := m.ctrl.Call(m, "IsValid", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid
func (mr *MockHasherMockRecorder) IsValid(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockHasher)(nil).IsValid), arg0, arg1)
}
