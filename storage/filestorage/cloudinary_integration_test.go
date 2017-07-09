package filestorage_test

import (
	"os"
	"testing"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
)

func TestCloudinary(t *testing.T) {
	if os.Getenv("TEST_CLOUDINARY") != "true" {
		t.Skip("Not testing cloudinary")
	}

	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	secret := os.Getenv("CLOUDINARY_SECRET")
	bucket := os.Getenv("CLOUDINARY_BUCKET")

	storage := filestorage.NewCloudinary(apiKey, secret)
	storage.SetBucket(bucket)

	storageHappyPathTest(t, storage, true)
	storageUnexistingReadTest(t, storage)
}
