package guard_test

import (
	"net/url"
	"testing"

	"github.com/Nivl/go-rest-tools/router/guard"
	"github.com/Nivl/go-rest-tools/security/auth"
	"github.com/stretchr/testify/assert"
)

type BasicParamStruct struct {
	UUID     string `from:"url" json:"uuid"  params:"uuid"`
	Trimmed  string `from:"form" json:"trimmed" params:"trim"`
	Required string `from:"form" json:"required" params:"required"`
	Pointer  *int   `from:"form" json:"pointer"`
	WontTrim string `from:"query" json:"wont_trim"`
}

// TestParseParams only test that the struct gets parsed and that data get
// returned, as well that it fails when it's supposed to.
// It won't check that the returned data are valid as it's
// out of this package scope.
func TestParseParams(t *testing.T) {
	shouldFail := true

	testCases := []struct {
		description string
		shouldFail  bool
		sources     map[string]url.Values
	}{
		{
			"valid complete struct should parse",
			!shouldFail,
			map[string]url.Values{
				"url": url.Values{
					"uuid": []string{"5847a692-88d3-4c1f-aedf-62622520d128"},
				},
				"form": url.Values{
					"trimmed":  []string{"   trimmed data   "},
					"required": []string{"required data"},
					"pointer":  []string{"8"},
				},
				"query": url.Values{
					"wont_trim": []string{"   not trimmed data   "},
				},
			},
		},
		{
			"Nil pointer should parse",
			!shouldFail,
			map[string]url.Values{
				"url": url.Values{
					"uuid": []string{"5847a692-88d3-4c1f-aedf-62622520d128"},
				},
				"form": url.Values{
					"trimmed":  []string{"   trimmed data   "},
					"required": []string{"required data"},
				},
				"query": url.Values{
					"wont_trim": []string{"   not trimmed data   "},
				},
			},
		},
		{
			"Missing required should failed to parse",
			shouldFail,
			map[string]url.Values{
				"url": url.Values{
					"uuid": []string{"5847a692-88d3-4c1f-aedf-62622520d128"},
				},
				"form": url.Values{
					"trimmed": []string{"   trimmed data   "},
				},
				"query": url.Values{
					"wont_trim": []string{"   not trimmed data   "},
				},
			},
		},
		{
			"invalid uuid should failed to parse",
			shouldFail,
			map[string]url.Values{
				"url": url.Values{
					"uuid": []string{"5847a692"},
				},
				"form": url.Values{
					"trimmed":  []string{"   trimmed data   "},
					"required": []string{"required data"},
				},
				"query": url.Values{
					"wont_trim": []string{"   not trimmed data   "},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			g := &guard.Guard{
				ParamStruct: &BasicParamStruct{},
			}

			data, err := g.ParseParams(tc.sources, nil)
			if tc.shouldFail {
				assert.NotNil(t, err, "the parsing was expected to fail")
				assert.Nil(t, data, "no data were expected to be returned")
			} else {
				assert.Nil(t, err, "the parsing was not expected to fail")
				assert.NotNil(t, data, "data were expected to be returned")
			}
		})
	}
}

func TestNoParseParams(t *testing.T) {
	testCases := []struct {
		description string
		guard       *guard.Guard
	}{
		{"no guards should work", nil},
		{"no paramStruct should work", &guard.Guard{}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			data, err := tc.guard.ParseParams(map[string]url.Values{}, nil)
			assert.Nil(t, err, "the parsing was not expected to fail")
			assert.Nil(t, data, "no data were expected to be returned")
		})
	}
}

// TestAuth just tests that it globally works and that the returned data
// are synced. It won't test all the RouteAuth functions as they already have
// their own tests.
func TestAuth(t *testing.T) {
	// sugar to avoid hardcoding true/false in all tests
	shouldFail := true

	testCases := []struct {
		description string
		guard       *guard.Guard
		user        *auth.User
		shouldFail  bool
	}{
		{"no guards should work", nil, nil, !shouldFail},
		{"no Auth should work", &guard.Guard{}, nil, !shouldFail},
		{
			"valid auth",
			&guard.Guard{Auth: guard.LoggedUserAccess},
			&auth.User{ID: "xxx"},
			!shouldFail,
		},
		{
			"invalid auth",
			&guard.Guard{Auth: guard.LoggedUserAccess},
			nil,
			shouldFail,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			success, err := tc.guard.HasAccess(tc.user)
			if tc.shouldFail {
				assert.False(t, success, "the access should have been denied")
				assert.NotNil(t, err, "an error should have been returned")
			} else {
				assert.True(t, success, "the access should have been granted")
				assert.Nil(t, err, "no error should have been returned")
			}
		})
	}
}
