package filestorage_test

import (
	"testing"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
)

func TestFSStorage(t *testing.T) {
	storage, err := filestorage.NewFSStorage()
	if err != nil {
		t.Fatal(err)
	}
	storage.SetBucket("unit-test")

	// we skip ValidateURL as the test uses http.Get which won't work with a
	// local file. Also Exists() calls URL() and will basically do the exact same
	// test
	cbs := &storageHappyPathTestCallbacks{
		ValidateURL: func(_ *testing.T, _ string) {},
	}
	storageHappyPathTest(t, storage, cbs)
	storageUnexistingReadTest(t, storage)
}
