package reporter_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dchest/uniuri"

	"github.com/satori/go.uuid"

	"github.com/Nivl/go-rest-tools/security/auth/testauth"

	"github.com/stretchr/testify/require"

	"github.com/Nivl/go-rest-tools/notifiers/reporter"
)

func TestSentryHappyPath(t *testing.T) {
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		t.Skip("sentry not set")
	}

	creator, err := reporter.NewSentryCreator(sentryDSN)
	require.NoError(t, err, "NewSentryCreator() should not have failed")

	r, err := creator.New()
	require.NoError(t, err, "creator.New() should not have failed")

	// Set some data
	r.SetUser(testauth.NewUser())
	r.AddTag("endpoint", "TestSentryHappyPath")
	r.AddTag("Req ID", uuid.NewV4().String())

	// send the request
	r.ReportErrorAndWait(fmt.Errorf("New test error %s", uniuri.NewLen(4)))
}
