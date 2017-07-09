package filestorage_test

import (
	"context"
	"os"
	"testing"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
)

func TestGCStorageUploadHappyPath(t *testing.T) {
	if os.Getenv("TEST_GCP") != "true" {
		t.Skip("Not testing GCP")
	}

	apiKey := os.Getenv("GCP_API_KEY")
	bucket := os.Getenv("GCP_BUCKET")

	ctx := context.Background()
	storage, err := filestorage.NewGCStorage(ctx, apiKey)
	if err != nil {
		t.Fatal(err)
	}
	storage.SetBucket(bucket)
	storageHappyPathTest(t, storage, true)
	storageUnexistingReadTest(t, storage)
}
