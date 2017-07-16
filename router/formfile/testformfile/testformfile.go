package testformfile

import (
	"mime/multipart"
	"os"
	"path"
	"testing"
)

// NewMultipartData is a helper to generate multipart data that can be returned
// by FileHolder.FormFile()
func NewMultipartData(t *testing.T, cwd string, filename string) (*multipart.FileHeader, *os.File) {
	// We create the Head and the file
	licenseHeader := &multipart.FileHeader{
		Filename: filename,
	}
	filePath := path.Join(cwd, "fixtures", filename)
	licenseFile, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	return licenseHeader, licenseFile
	//defer licenseFile.Close()
}
