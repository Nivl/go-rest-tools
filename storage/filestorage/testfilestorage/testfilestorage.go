package testfilestorage

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
	"github.com/stretchr/testify/assert"
)

// StorageHappyPathTestCallbacks is a structure to overide/skip certain test
// It's useful when a specific provider react differently than the expected
// behavior (like deleted file on Cloudinary that are still available for
// a few hours)
type StorageHappyPathTestCallbacks struct {
	ValidateURL              func(*testing.T, string)
	StillExistsAfterDeletion func(*testing.T, string)
}

// StorageHappyPathTest is a helper to test the happy path of FileStorage
func StorageHappyPathTest(t *testing.T, storage filestorage.FileStorage, callbacks *StorageHappyPathTestCallbacks) {
	if callbacks == nil {
		callbacks = &StorageHappyPathTestCallbacks{}
	}

	testCases := []struct {
		description string
		outputName  string
		fileContent []byte
	}{
		{
			"Subfolder should work",
			"subfolder/" + strconv.FormatInt(time.Now().UTC().Unix(), 10),
			[]byte("this is the content of the testfile"),
		},
		{
			"No subfolder should work",
			"no-subfolder-" + strconv.FormatInt(time.Now().UTC().Unix(), 10),
			[]byte("this is the content of the other testfile"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			reader := bytes.NewReader(tc.fileContent)

			// Upload the file
			if err := storage.Write(reader, tc.outputName); err != nil {
				assert.FailNow(t, err.Error(), "expected Write() to succeed")
			}

			// Make sure the file exists
			exists, err := storage.Exists(tc.outputName)
			if err != nil {
				assert.FailNow(t, err.Error(), "expected Exists() to succeed")
			}
			assert.True(t, exists, "expected the file to exist")

			// Read the uploaded file
			r, err := storage.Read(tc.outputName)
			if err != nil {
				assert.FailNow(t, err.Error(), "expected Read() to succeed")
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(r)
			r.Close()
			assert.Equal(t, tc.fileContent, buf.Bytes(), "the file content shouldn't have changed")

			// Make sure a URL is valid
			if callbacks.ValidateURL != nil {
				callbacks.ValidateURL(t, tc.outputName)
			} else {
				url, err := storage.URL(tc.outputName)
				if err != nil {
					assert.FailNow(t, err.Error(), "couldn't find the URL of a file")
				}
				resp, err := http.Get(url)
				if err != nil {
					assert.FailNow(t, err.Error(), "couldn't GET the URL of a file")
				}
				resp.Body.Close()
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			}

			// Delete the file
			err = storage.Delete(tc.outputName)
			assert.NoError(t, err, "the deletion should have succeed")

			// Make sure the file is deleted
			if callbacks.StillExistsAfterDeletion != nil {
				callbacks.StillExistsAfterDeletion(t, tc.outputName)
			} else {
				exists, err = storage.Exists(tc.outputName)
				assert.NoError(t, err, "expect Exists() to succeed for unexisting file")
				assert.False(t, exists, "expect the file not to exist")
			}
		})
	}
}

// StorageUnexistingReadTest tests that
func StorageUnexistingReadTest(t *testing.T, storage filestorage.FileStorage) {
	testCases := []struct {
		description string
		filename    string
	}{
		{
			"Subfolder should work",
			"subfolder/file-does-not-exist",
		},
		{
			"No subfolder should work",
			"file-does-not-exist",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// Read the uploaded file
			r, err := storage.Read(tc.filename)
			if err == nil {
				r.Close()
			}

			assert.Error(t, err, "expected Read() to fail")
			assert.True(t, os.IsNotExist(err), "expected Read() to fail with a Not Found error")
		})
	}
}
