package router

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseJSONBody(t *testing.T) {
	body := `{"array":[1,2,3], "string": "value"}`

	req := &HTTPRequest{
		http: &http.Request{
			Body: ioutil.NopCloser(strings.NewReader(body)),
			Header: http.Header{
				"Content-Type": []string{ContentTypeJSON},
			},
		},
	}

	vals, err := req.parseJSONBody()
	require.NoError(t, err, "parseJSONBody() should have succeed")
	assert.Equal(t, vals["array"], []string{"1", "2", "3"}, "invalid values for 'array'")
	assert.Equal(t, []string{"value"}, vals["string"], "invalid value for 'string'")
}
