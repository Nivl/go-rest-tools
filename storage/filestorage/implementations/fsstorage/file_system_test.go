package fsstorage_test

import (
	"testing"

	"github.com/Nivl/go-rest-tools/storage/filestorage/implementations/fsstorage"
	"github.com/Nivl/go-rest-tools/storage/filestorage/testfilestorage"
)

func TestFSStorage(t *testing.T) {
	storage, err := fsstorage.New()
	if err != nil {
		t.Fatal(err)
	}
	storage.SetBucket("unit-test")

	// we skip ValidateURL as the test uses http.Get which won't work with a
	// local file. Also Exists() calls URL() and will basically do the exact same
	// test
	cbs := &testfilestorage.StorageHappyPathTestCallbacks{
		ValidateURL: func(_ *testing.T, _ string) {},
	}
	testfilestorage.StorageHappyPathTest(t, storage, cbs)
	testfilestorage.StorageUnexistingReadTest(t, storage)
}
