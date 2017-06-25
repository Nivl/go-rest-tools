package guard_test

import (
	"net/http"
	"testing"

	"github.com/Nivl/go-rest-tools/network/http/httperr"
	"github.com/Nivl/go-rest-tools/primitives/ptrs"
	"github.com/Nivl/go-rest-tools/router/guard"
	"github.com/Nivl/go-rest-tools/security/auth"
	"github.com/stretchr/testify/assert"
)

func TestLoggedUserAccess(t *testing.T) {
	testCases := []struct {
		description   string
		user          *auth.User
		expectedError *int
	}{
		{
			"Anonymous user",
			nil,
			ptrs.NewInt(http.StatusUnauthorized),
		},
		{
			"Invalid user object",
			&auth.User{},
			ptrs.NewInt(http.StatusUnauthorized),
		},
		{
			"Logged In user",
			&auth.User{ID: "xxx"},
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			err := guard.LoggedUserAccess(tc.user)
			if tc.expectedError == nil {
				assert.Nil(t, err, "access should have not been denied: %s", err)
			} else {
				errCode := getHTTPError(err)
				assert.Equal(t, *tc.expectedError, errCode, "the auth failed with the wrong error code")
			}
		})
	}
}

func TestAdminAccess(t *testing.T) {
	testCases := []struct {
		description   string
		user          *auth.User
		expectedError *int
	}{
		{
			"Anonymous user",
			nil,
			ptrs.NewInt(http.StatusUnauthorized),
		},
		{
			"Logged In user",
			&auth.User{ID: "xxx"},
			ptrs.NewInt(http.StatusForbidden),
		},
		{
			"Admin",
			&auth.User{ID: "xxx", IsAdmin: true},
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			err := guard.AdminAccess(tc.user)
			if tc.expectedError == nil {
				assert.Nil(t, err, "access should have not been denied: %s", err)
			} else {
				errCode := getHTTPError(err)
				assert.Equal(t, *tc.expectedError, errCode, "the auth failed with the wrong error code")
			}
		})
	}
}

// Helpers

func getHTTPError(e error) int {
	err, casted := e.(*httperr.HTTPError)
	if !casted {
		return http.StatusInternalServerError
	}
	return err.Code()
}
