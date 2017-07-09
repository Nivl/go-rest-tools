package filestorage_test

import (
	"bytes"
	"os"
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
			url, _ := fs.URL(tc.outputName)
			_, err := os.Stat(url)
			if err != nil {
				if os.IsNotExist(err) {
					assert.FailNow(t, "Expected the following file to exists: "+url)
				}
				t.Fatal(err)
			}

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
			_, err = os.Stat(url)
			assert.True(t, os.IsNotExist(err), "expect the file not to exist")
		})
	}
}
