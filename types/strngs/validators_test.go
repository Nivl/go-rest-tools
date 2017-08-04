package strngs_test

import (
	"testing"

	"github.com/Nivl/go-rest-tools/types/strngs"
	"github.com/stretchr/testify/assert"
)

func TestIsValidURL(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description string
		uri         string
		shouldFail  bool
	}{
		{"ftp should fail", "ftp://google.com", shouldFail},
		{"file should fail", "file:///dev/urandom", shouldFail},
		{"http should work", "http://google.com", !shouldFail},
		{"https should work", "https://google.com", !shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, !tc.shouldFail, strngs.IsValidURL(tc.uri))
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description string
		email       string
		shouldFail  bool
	}{
		{"valid email should work", "email@domain.tld", !shouldFail},
		{"email with a + should work", "email+filter@domain.tld", !shouldFail},
		{"email without @ should fail", "emaildomain.tld", shouldFail},
		{"email without . should fail", "email@domaintld", shouldFail},
		{"email with nothing before @ should fail", "@domain.tld", shouldFail},
		{"email with nothing after . should fail", "email@domain.", shouldFail},
		{"email with nothing before . should fail", "email@.tld", shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, !tc.shouldFail, strngs.IsValidEmail(tc.email))
		})
	}
}
