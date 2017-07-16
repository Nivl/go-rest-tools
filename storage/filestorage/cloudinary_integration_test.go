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

	// we skip StillExistsAfterDeletion because Cloudinary doesn't remove the
	// files right away, so this test will always fail
	cbs := &storageHappyPathTestCallbacks{
		StillExistsAfterDeletion: func(_ *testing.T, _ string) {},
	}
	storageHappyPathTest(t, storage, cbs)
	storageUnexistingReadTest(t, storage)
}
