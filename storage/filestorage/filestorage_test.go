package filestorage_test

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

// storageHappyPathTest is a helper to test the happy path of FileStorage
// if dontTestIfRemoved is set to true, the function won't check if the
// uploaded is still accessible after the deletion. This is useful for
// providers like Cloudinary that can take an hour to delete a file
func storageHappyPathTest(t *testing.T, storage filestorage.FileStorage, dontTestIfRemoved bool) {
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

			// Delete the file
			err = storage.Delete(tc.outputName)
			assert.NoError(t, err, "the deletion should have succeed")

			if !dontTestIfRemoved {
				//Make sure the file is deleted
				r, err = storage.Read(tc.outputName)
				assert.Error(t, err, "Read should have fail")
				if err == nil {
					r.Close()
				}
				assert.Equal(t, os.ErrNotExist, err, "it should have failed because the file no longer exist")
			}
		})
	}
}
