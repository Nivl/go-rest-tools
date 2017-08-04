package paginator_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Nivl/go-rest-tools/paginator"
	"github.com/Nivl/go-rest-tools/primitives/apierror"
	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaginator(t *testing.T) {
	testCases := []struct {
		description    string
		p              *paginator.Paginator
		shouldBeValid  bool
		expectedOffset int
		expectedLimit  int
	}{
		{
			"Page 1, PerPage 100",
			paginator.New(1, 100),
			true,
			0,
			100,
		},
		{
			"Page 6, PerPage 600",
			paginator.New(6, 600),
			false,
			0, 0,
		},
		{
			"Page 50, PerPage 10",
			paginator.New(50, 10),
			true,
			490,
			10,
		},
		{
			"Page 0, PerPage 10",
			paginator.New(0, 10),
			false, 0, 0,
		},
		{
			"Page 10, PerPage 0",
			paginator.New(10, 0),
			false, 0, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.shouldBeValid, tc.p.IsValid())
			if tc.shouldBeValid {
				assert.Equal(t, tc.expectedLimit, tc.p.Limit())
				assert.Equal(t, tc.expectedOffset, tc.p.Offset())
			}
		})
	}
}

func TestHandlerParams(t *testing.T) {
	// sugar
	shouldFail := true

	type strct struct {
		paginator.HandlerParams
	}

	testCases := []struct {
		description         string
		params              url.Values
		shouldFail          bool
		expectedErrorField  string
		expectedErrorMsg    string
		expectedCurrentPage int
		expectedPerPage     int
	}{
		{
			"Default values",
			url.Values{},
			!shouldFail,
			"", "",
			1, 100,
		},
		{
			"Page 4, PerPage 50",
			url.Values{
				"page":     []string{"4"},
				"per_page": []string{"50"},
			},
			!shouldFail,
			"", "",
			4, 50,
		},
		{
			"PerPage 8",
			url.Values{
				"per_page": []string{"8"},
			},
			!shouldFail,
			"", "",
			1, 8,
		},
		{
			"Page 40",
			url.Values{
				"page": []string{"40"},
			},
			!shouldFail,
			"", "",
			40, 100,
		},
		{
			"page set to 0 should fail",
			url.Values{
				"page": []string{"0"},
			},
			shouldFail,
			"page", "cannot be <= 0",
			0, 0,
		},
		{
			"per_page set to 0 should fail",
			url.Values{
				"per_page": []string{"0"},
			},
			shouldFail,
			"per_page", "cannot be <= 0",
			0, 0,
		},
		{
			"per_page above 100 should fail",
			url.Values{
				"per_page": []string{"101"},
			},
			shouldFail,
			"per_page", "cannot be > 100",
			0, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			s := &strct{}
			p := params.NewParams(s)
			sources := map[string]url.Values{
				"query": tc.params,
			}

			err := p.Parse(sources, nil)
			if tc.shouldFail {
				require.Error(t, err, "Parse() should have failed")

				e := apierror.Convert(err)
				assert.Equal(t, http.StatusBadRequest, e.HTTPStatus(), "It should have failed with a 400")
				assert.Equal(t, tc.expectedErrorField, e.Field(), "Failed on the wrong field")
				assert.True(t, strings.Contains(err.Error(), tc.expectedErrorMsg),
					"the error \"%s\" should contain the string \"%s\"", err.Error(), tc.expectedErrorMsg)
			} else {
				assert.NoError(t, err, "Parse() should have succeed")
				assert.Equal(t, tc.expectedCurrentPage, s.Page)
				assert.Equal(t, tc.expectedPerPage, s.PerPage)
			}
		})
	}
}
