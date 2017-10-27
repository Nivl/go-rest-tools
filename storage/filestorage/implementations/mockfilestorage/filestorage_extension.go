package mockfilestorage

import (
	"errors"
	"os"
	"path"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
	gomock "github.com/golang/mock/gomock"
)

var (
	// StringType represent a string argument
	StringType = gomock.Eq("string")

	// UpdatableAttrType represents a *filestorage.UpdatableFileAttributes argument
	UpdatableAttrType = gomock.Eq("*filestorage.UpdatableFileAttributes")

	// AnyType represents am argument that can accept anything
	AnyType = gomock.Any()
)

// ExpectWriteIfNotExist is an helper that expects a WriteIfNotExist to
// succeed with the provided params
func (s *MockFileStorage) ExpectWriteIfNotExist(isNew bool, url string) *gomock.Call {
	return s.EXPECT().WriteIfNotExist(AnyType, StringType).Return(isNew, url, nil)
}

// ExpectWriteIfNotExistError is an helper that expects a WriteIfNotExist to
// fail
func (s *MockFileStorage) ExpectWriteIfNotExistError() *gomock.Call {
	call := s.EXPECT().WriteIfNotExist(AnyType, StringType)
	call.Return(false, "", errors.New("server unreachable"))
	return call
}

// ExpectRead is an helper that expects a Read
func (s *MockFileStorage) ExpectRead(cwd, filename string) *gomock.Call {
	filePath := path.Join(cwd, "fixtures", filename)
	return s.EXPECT().Read(StringType).Return(os.Open(filePath))
}

// ExpectExists is an helper that expects Exists() to return true
func (s *MockFileStorage) ExpectExists() *gomock.Call {
	return s.EXPECT().Exists(StringType).Return(true, nil)
}

// ExpectNotExists is an helper that expects Exists() to return false
func (s *MockFileStorage) ExpectNotExists() *gomock.Call {
	return s.EXPECT().Exists(StringType).Return(false, nil)
}

// ExpectURL is an helper that expects URL() to return given param
func (s *MockFileStorage) ExpectURL(url string) *gomock.Call {
	return s.EXPECT().URL(StringType).Return(url, nil)
}

// ExpectSetAttributes is an helper that expects SetAttributes() to work,
// and returns an empty content
func (s *MockFileStorage) ExpectSetAttributes() *gomock.Call {
	attrs := &filestorage.FileAttributes{}
	return s.EXPECT().SetAttributes(StringType, UpdatableAttrType).Return(attrs, nil)
}

// ExpectSetAttributesRet is an helper that expects SetAttributes() to
// return the provided object
func (s *MockFileStorage) ExpectSetAttributesRet(attrs *filestorage.FileAttributes) *gomock.Call {
	return s.EXPECT().SetAttributes(StringType, UpdatableAttrType).Return(attrs, nil)
}

// ExpectAttributes is an helper that expects Attributes() to
// return the provided object
func (s *MockFileStorage) ExpectAttributes(attrs *filestorage.FileAttributes) *gomock.Call {
	return s.EXPECT().Attributes(StringType).Return(attrs, nil)
}

// ExpectDelete is an helper that expects Delete() to succeed
func (s *MockFileStorage) ExpectDelete() *gomock.Call {
	return s.EXPECT().Delete(StringType).Return(nil)
}
