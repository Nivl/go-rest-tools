package cloudinary

import (
	"context"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
)

// Makes sure Creator is a logger.Creator
var _ filestorage.Creator = (*Creator)(nil)

// NewCreator returns a filestorage creator that will use the provided keys
// to create a new cloudinary driver for each single logger
func NewCreator(apiKey, secret string) *Creator {
	return &Creator{
		apiKey: apiKey,
		secret: secret,
	}
}

// Creator creates new filestorage
type Creator struct {
	apiKey string
	secret string
}

// New returns a new le client
func (c *Creator) New(ctx context.Context) (filestorage.FileStorage, error) {
	return New(c.apiKey, c.secret), nil
}
