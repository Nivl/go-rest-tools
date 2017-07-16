package filestorage_test

import (
	"bytes"
	"testing"

	"strconv"
	"time"

	"github.com/Nivl/go-rest-tools/storage/filestorage"
	"github.com/stretchr/testify/assert"
)

func TestFSStorage(t *testing.T) {
	storage, err := filestorage.NewFSStorage()
	if err != nil {
		t.Fatal(err)
	}
	storage.SetBucket("unit-test")

	fsHappyPath(t, storage)
	storageUnexistingReadTest(t, storage)
}

func fsHappyPath(t *testing.T, fs filestorage.FileStorage) {
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
			if err := fs.Write(reader, tc.outputName); err != nil {
				assert.FailNow(t, err.Error(), "expected Write() to succeed")
			}

			// check the file exist on the disk
			exists, err := fs.Exists(tc.outputName)
			if err != nil {
				assert.FailNow(t, err.Error(), "expected Exists() to succeed")
			}
			assert.True(t, exists, "expected the file to exist")

			// Read the uploaded file
			r, err := fs.Read(tc.outputName)
			if err != nil {
				assert.FailNow(t, err.Error(), "expected Read() to succeed")
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(r)
			r.Close()
			assert.Equal(t, tc.fileContent, buf.Bytes(), "the file content shouldn't have changed")

			// Delete the file
			err = fs.Delete(tc.outputName)
			assert.NoError(t, err, "the deletion should have succeed")

			// We make sure the file has been deleted
			exists, err = fs.Exists(tc.outputName)
			assert.NoError(t, err, "expect Exists() to succeed for unexisting file")
			assert.False(t, exists, "expect the file not to exist")
		})
	}
}
