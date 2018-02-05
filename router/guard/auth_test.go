package guard_test

import (
	"testing"

	"github.com/Nivl/go-rest-tools/types/apperror"

	"github.com/Nivl/go-rest-tools/router/guard"
	"github.com/Nivl/go-rest-tools/security/auth"
	"github.com/Nivl/go-types/ptrs"
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
			ptrs.NewInt(int(apperror.Unauthenticated)),
		},
		{
			"Invalid user object",
			&auth.User{},
			ptrs.NewInt(int(apperror.Unauthenticated)),
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
				assert.Equal(t, *tc.expectedError, int(err.StatusCode()), "the auth failed with the wrong error code")
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
			ptrs.NewInt(int(apperror.Unauthenticated)),
		},
		{
			"Logged In user",
			&auth.User{ID: "xxx"},
			ptrs.NewInt(int(apperror.PermissionDenied)),
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
				assert.Equal(t, *tc.expectedError, int(err.StatusCode()), "the auth failed with the wrong error code")
			}
		})
	}
}
