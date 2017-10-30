// Package gcstorage is an implementation of filestorage for Google Cloud
package gcstorage

import (
	"context"
	"io"

	"github.com/Nivl/go-filestorage"
	"github.com/Nivl/go-filestorage/implementations"
	"google.golang.org/api/option"

	"fmt"

	"cloud.google.com/go/storage"
)

// updatableAttrsToGCStorage converts a *UpdatableFileAttributes into a *storage.ObjectAttrsToUpdate
func updatableAttrsToGCStorage(attrs *filestorage.UpdatableFileAttributes) *storage.ObjectAttrsToUpdate {
	return &storage.ObjectAttrsToUpdate{
		ContentType:        attrs.ContentType,
		ContentDisposition: attrs.ContentDisposition,
		ContentLanguage:    attrs.ContentLanguage,
		ContentEncoding:    attrs.ContentEncoding,
		CacheControl:       attrs.CacheControl,
		Metadata:           attrs.Metadata,
	}
}

// NewAttributes converts a *storage.ObjectAttrs to *FileAttributes
func NewAttributes(attrs *storage.ObjectAttrs) *filestorage.FileAttributes {
	return &filestorage.FileAttributes{
		ContentType:        attrs.ContentType,
		ContentDisposition: attrs.ContentDisposition,
		ContentLanguage:    attrs.ContentLanguage,
		ContentEncoding:    attrs.ContentEncoding,
		CacheControl:       attrs.CacheControl,
		Metadata:           attrs.Metadata,
	}
}

// New returns a new instance of a Google Cloud Storage
func New(ctx context.Context, apiKey string) (*GCStorage, error) {
	client, err := storage.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &GCStorage{
		ctx:    ctx,
		client: client,
	}, nil
}

// GCStorage is an implementation of the FileStorage interface for Google Cloud
type GCStorage struct {
	ctx        context.Context
	client     *storage.Client
	bucket     *storage.BucketHandle
	bucketName string
}

// ID returns the unique identifier of the storage provider
func (s *GCStorage) ID() string {
	return "google_cloud_storage"
}

// SetBucket is used to set the bucket
// Always return nil
func (s *GCStorage) SetBucket(name string) error {
	s.bucketName = name
	s.bucket = s.client.Bucket(name)
	return nil
}

// Read fetches a file a returns a reader
func (s *GCStorage) Read(filepath string) (io.ReadCloser, error) {
	return s.bucket.Object(filepath).NewReader(s.ctx)
}

// Write copy the provided os.File to dest
func (s *GCStorage) Write(src io.Reader, destPath string) error {
	obj := s.bucket.Object(destPath)
	dest := obj.NewWriter(s.ctx)

	// Copy the file
	_, err := io.Copy(dest, src)
	if err != nil {
		dest.CloseWithError(err)
		return err
	}

	// Send the changes to GCP
	return dest.Close()
}

// SetAttributes sets the attributes of the file
func (s *GCStorage) SetAttributes(filepath string, attrs *filestorage.UpdatableFileAttributes) (*filestorage.FileAttributes, error) {
	gcsAttrs, err := s.bucket.Object(filepath).Update(s.ctx, *updatableAttrsToGCStorage(attrs))
	if err != nil {
		return nil, err
	}

	return NewAttributes(gcsAttrs), nil
}

// Attributes returns the attributes of the file
func (s *GCStorage) Attributes(filepath string) (*filestorage.FileAttributes, error) {
	gcsAttrs, err := s.bucket.Object(filepath).Attrs(s.ctx)
	if err != nil {
		return nil, err
	}
	return NewAttributes(gcsAttrs), nil
}

// Exists check if a file exists
func (s *GCStorage) Exists(filepath string) (bool, error) {
	_, err := s.Attributes(filepath)
	if err == nil {
		return true, nil
	}
	if err == storage.ErrObjectNotExist {
		return false, nil
	}
	return false, err
}

// URL returns the URL of the file
func (s *GCStorage) URL(filepath string) (string, error) {
	return fmt.Sprintf("https://%s.storage.googleapis.com/%s", s.bucketName, filepath), nil
}

// Delete removes a file, ignores files that do not exist
func (s *GCStorage) Delete(filepath string) error {
	return s.bucket.Object(filepath).Delete(s.ctx)
}

// WriteIfNotExist copies the provided io.Reader to dest if the file does
// not already exist
// Returns:
//   - A boolean specifying if the file got uploaded (true) or if already
//     existed (false).
//   - A URL to the uploaded file
//   - An error if something went wrong
func (s *GCStorage) WriteIfNotExist(src io.Reader, destPath string) (new bool, url string, err error) {
	return implementations.WriteIfNotExist(s, src, destPath)
}
