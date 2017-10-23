package cloudinary_test

import (
	"os"
	"testing"

	"github.com/Nivl/go-rest-tools/storage/filestorage/implementations/cloudinary"
	"github.com/Nivl/go-rest-tools/storage/filestorage/testfilestorage"
)

func TestCloudinary(t *testing.T) {
	if os.Getenv("TEST_CLOUDINARY") != "true" {
		t.Skip("Not testing cloudinary")
	}

	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	secret := os.Getenv("CLOUDINARY_SECRET")
	bucket := os.Getenv("CLOUDINARY_BUCKET")

	storage := cloudinary.New(apiKey, secret)
	storage.SetBucket(bucket)

	// we skip StillExistsAfterDeletion because Cloudinary doesn't remove the
	// files right away, so this test will always fail
	cbs := &testfilestorage.StorageHappyPathTestCallbacks{
		StillExistsAfterDeletion: func(_ *testing.T, _ string) {},
	}
	testfilestorage.StorageHappyPathTest(t, storage, cbs)
	testfilestorage.StorageUnexistingReadTest(t, storage)
}
