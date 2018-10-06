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
	t.Run("valid data", func(t *testing.T) {
		body := `{"array":[1,2,3], "string": "value", "empty_array":[]}`

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
		assert.Equal(t, []string{"1", "2", "3"}, vals["array"], "invalid values for 'array'")
		assert.Equal(t, []string{"value"}, vals["string"], "invalid value for 'string'")
		assert.Equal(t, []string{}, vals["empty_array"], "invalid value for 'empty_array'")
	})

	t.Run("invalid JSON data", func(t *testing.T) {
		body := "{"

		req := &HTTPRequest{
			http: &http.Request{
				Body: ioutil.NopCloser(strings.NewReader(body)),
				Header: http.Header{
					"Content-Type": []string{ContentTypeJSON},
				},
			},
		}

		_, err := req.parseJSONBody()
		require.Error(t, err, "parseJSONBody() should have failed")
		require.Equal(t, ErrMsgInvalidJSONPayload, err.Error(), "unexpected error returned")
	})
}

func TestContentType(t *testing.T) {
	t.Run("json utf8", func(t *testing.T) {
		req := &HTTPRequest{
			http: &http.Request{
				Header: http.Header{
					"Content-Type": []string{ContentTypeJSON + "; charset=utf-8"},
				},
			},
		}

		ct := req.contentType()
		assert.Equal(t, ContentTypeJSON, ct, "invalid content type")
	})

	t.Run("basic json", func(t *testing.T) {
		req := &HTTPRequest{
			http: &http.Request{
				Header: http.Header{
					"Content-Type": []string{ContentTypeJSON},
				},
			},
		}

		ct := req.contentType()
		assert.Equal(t, ContentTypeJSON, ct, "invalid content type")
	})

	t.Run("empty", func(t *testing.T) {
		req := &HTTPRequest{
			http: &http.Request{
				Header: http.Header{
					"Content-Type": []string{""},
				},
			},
		}

		ct := req.contentType()
		assert.Empty(t, ct, "invalid content type")
	})
}
