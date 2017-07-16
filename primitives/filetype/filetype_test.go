package filetype_test

import (
	"bytes"
	"testing"

	"github.com/Nivl/go-rest-tools/primitives/filetype"
	"github.com/stretchr/testify/assert"
)

func TestSHA256Sum(t *testing.T) {
	testCases := []struct {
		content  string
		expected string
	}{
		{"this is a test", "2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.content, func(t *testing.T) {
			t.Parallel()

			r := bytes.NewReader([]byte(tc.content))
			sum, err := filetype.SHA256Sum(r)
			assert.NoError(t, err, "SHA256Sum() should have succeed")
			assert.Equal(t, tc.expected, sum, "invalid sum")
		})
	}
}
