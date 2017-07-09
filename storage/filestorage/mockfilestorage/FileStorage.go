// Code generated by mockery v1.0.0
package mockfilestorage

import filestorage "github.com/Nivl/go-rest-tools/storage/filestorage"
import io "io"
import mock "github.com/stretchr/testify/mock"

// FileStorage is an autogenerated mock type for the FileStorage type
type FileStorage struct {
	mock.Mock
}

// Attributes provides a mock function with given fields: filepath
func (_m *FileStorage) Attributes(filepath string) (*filestorage.FileAttributes, error) {
	ret := _m.Called(filepath)

	var r0 *filestorage.FileAttributes
	if rf, ok := ret.Get(0).(func(string) *filestorage.FileAttributes); ok {
		r0 = rf(filepath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*filestorage.FileAttributes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filepath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: filepath
func (_m *FileStorage) Delete(filepath string) error {
	ret := _m.Called(filepath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filepath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *FileStorage) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Read provides a mock function with given fields: filepath
func (_m *FileStorage) Read(filepath string) (io.ReadCloser, error) {
	ret := _m.Called(filepath)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(filepath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filepath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAttributes provides a mock function with given fields: filepath, attrs
func (_m *FileStorage) SetAttributes(filepath string, attrs *filestorage.UpdatableFileAttributes) (*filestorage.FileAttributes, error) {
	ret := _m.Called(filepath, attrs)

	var r0 *filestorage.FileAttributes
	if rf, ok := ret.Get(0).(func(string, *filestorage.UpdatableFileAttributes) *filestorage.FileAttributes); ok {
		r0 = rf(filepath, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*filestorage.FileAttributes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *filestorage.UpdatableFileAttributes) error); ok {
		r1 = rf(filepath, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetBucket provides a mock function with given fields: bucket
func (_m *FileStorage) SetBucket(bucket string) error {
	ret := _m.Called(bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// URL provides a mock function with given fields: filepath
func (_m *FileStorage) URL(filepath string) (string, error) {
	ret := _m.Called(filepath)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filepath)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filepath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: src, destPath
func (_m *FileStorage) Write(src io.Reader, destPath string) error {
	ret := _m.Called(src, destPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Reader, string) error); ok {
		r0 = rf(src, destPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
