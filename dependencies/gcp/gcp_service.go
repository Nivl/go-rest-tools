package gcp

import (
	"context"

	filestorage "github.com/Nivl/go-filestorage"
	"github.com/Nivl/go-filestorage/implementations/gcstorage"
)

var _ GCP = (*Service)(nil)

// New creates a new Service instance
func New(apiKey, projectName, bucket string) GCP {
	return &Service{
		apiKey:      apiKey,
		projectName: projectName,
		bucket:      bucket,
	}
}

// Service implements the GCP interface
type Service struct {
	apiKey      string
	projectName string
	bucket      string

	storage filestorage.FileStorage
}

// Storage returns a new storage instance
func (service *Service) Storage(ctx context.Context) (filestorage.FileStorage, error) {
	storage, err := gcstorage.NewWithContext(ctx, service.apiKey)
	if err != nil {
		return nil, err
	}
	storage.SetBucket(service.bucket)
	return storage, nil
}
