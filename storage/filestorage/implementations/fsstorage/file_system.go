package fsstorage

import (
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
	"github.com/Nivl/go-rest-tools/storage/fs"
)

// New returns a new instance of a File System Storage
func New() (*FSStorage, error) {
	tmpDir, err := ioutil.TempDir("", "storage")
	if err != nil {
		return nil, err
	}

	return &FSStorage{
		path: tmpDir,
	}, nil
}

// NewWithDir returns a new instance of a File System Storage with
//
func NewWithDir(path string) *FSStorage {
	return &FSStorage{
		path: path,
	}
}

// FSStorage is an implementation of the FileStorage interface for the file system
type FSStorage struct {
	path   string
	bucket string
}

// ID returns the unique identifier of the storage provider
func (s *FSStorage) ID() string {
	return "file_system"
}

// SetBucket is used to set the bucket
func (s *FSStorage) SetBucket(name string) error {
	s.bucket = name
	return nil
}

// Read fetches a file a returns a reader
func (s *FSStorage) Read(filepath string) (io.ReadCloser, error) {
	return os.Open(s.fullPath(filepath))
}

// Exists checks if a file exists
func (s *FSStorage) Exists(filepath string) (bool, error) {
	return fs.FileExists(s.fullPath(filepath))
}

// Write copy the provided os.File to dest
func (s *FSStorage) Write(src io.Reader, destPath string) error {
	fullPath := s.fullPath(destPath)

	// make sure the path exists
	if err := os.MkdirAll(path.Dir(fullPath), os.ModePerm); err != nil {
		return err
	}

	dest, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	// Copy the file
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a file
func (s *FSStorage) Delete(filepath string) error {
	err := os.Remove(s.fullPath(filepath))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// URL returns the URL of the file
func (s *FSStorage) URL(filepath string) (string, error) {
	return s.fullPath(filepath), nil
}

// SetAttributes sets the attributes of the file
func (s *FSStorage) SetAttributes(filepath string, attrs *filestorage.UpdatableFileAttributes) (*filestorage.FileAttributes, error) {
	return filestorage.NewFileAttributesFromUpdatable(attrs), nil
}

// Attributes returns the attributes of the file
// Always returns an empty struct as no attributes are kept for this FS
func (s *FSStorage) Attributes(filepath string) (*filestorage.FileAttributes, error) {
	return &filestorage.FileAttributes{}, nil
}

func (s *FSStorage) fullPath(filepath string) string {
	return path.Join(s.path, s.bucket, filepath)
}

// WriteIfNotExist copies the provided io.Reader to dest if the file does
// not already exist
// Returns:
//   - A boolean specifying if the file got uploaded (true) or if already
//     existed (false).
//   - A URL to the uploaded file
//   - An error if something went wrong
func (s *FSStorage) WriteIfNotExist(src io.Reader, destPath string) (new bool, url string, err error) {
	return filestorage.WriteIfNotExist(s, src, destPath)
}
