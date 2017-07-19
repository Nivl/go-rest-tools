package fs_test

import (
	"os"
	"path"
	"testing"

	"github.com/Nivl/go-rest-tools/storage/fs"
	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		description string
		path        string
		exists      bool
	}{
		{"Existing file", path.Join(cwd, "fs.go"), true},
		{"Not existing file", path.Join(cwd, "fs.js"), false},
		{"Existing dir", path.Join(cwd), true},
		{"Not existing dir", path.Join(cwd, "invalid"), false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			exists, err := fs.Exists(tc.path)
			assert.NoError(t, err, "Exists() should have succeed")
			assert.Equal(t, tc.exists, exists, "Exists() did not return the expected value")
		})
	}
}

func TestFileExists(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		description string
		path        string
		exists      bool
	}{
		{"Existing file", path.Join(cwd, "fs.go"), true},
		{"Not existing file", path.Join(cwd, "fs.js"), false},
		{"Existing dir", path.Join(cwd), false},
		{"Not existing dir", path.Join(cwd, "invalid"), false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			exists, err := fs.FileExists(tc.path)
			assert.NoError(t, err, "FileExists() should have succeed")
			assert.Equal(t, tc.exists, exists, "FileExists() did not return the expected value")
		})
	}
}

func TestDirExists(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		description string
		path        string
		exists      bool
	}{
		{"Existing file", path.Join(cwd, "fs.go"), false},
		{"Not existing file", path.Join(cwd, "fs.js"), false},
		{"Existing dir", path.Join(cwd), true},
		{"Not existing dir", path.Join(cwd, "invalid"), false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			exists, err := fs.DirExists(tc.path)
			assert.NoError(t, err, "DirExists() should have succeed")
			assert.Equal(t, tc.exists, exists, "DirExists() did not return the expected value")
		})
	}
}
