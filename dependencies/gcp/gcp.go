package gcp

import (
	"context"

	filestorage "github.com/Nivl/go-filestorage"
)

// GCP represents a Google Cloud Platform service
type GCP interface {
	// Storage creates a new storage attached to the provided context
	Storage(ctx context.Context) (filestorage.FileStorage, error)
}
