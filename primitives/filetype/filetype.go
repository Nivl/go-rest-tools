package filetype

import (
	"io"
	"net/http"
)

// FileValidator represents a function that can validate a file type
type FileValidator func(r io.ReadSeeker) (bool, error)

// MimeType returns the mimetype of a file
func MimeType(r io.ReadSeeker) (string, error) {
	// DetectContentType needs the first 512 bytes
	bytesNeeded := 512
	buff := make([]byte, bytesNeeded)
	n, err := r.Read(buff)
	if err != nil {
		return "", err
	}

	// we seek back to where we were like we didn't do anything
	if _, err := r.Seek(int64(-n), io.SeekCurrent); err != nil {
		return "", err
	}

	return http.DetectContentType(buff), nil
}

// SHA256 returns the SHA256 sum of a reader
func SHA256(r io.ReadSeeker) (string, error) {
	initialPos, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return "", err
	}

	// revert the pointer back to its original position
	_, err = r.Seek(initialPos, io.SeekStart)
	if err != nil {
		return "", err
	}

	return "", nil
}
