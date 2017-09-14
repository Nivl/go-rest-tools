package httptests_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Nivl/go-rest-tools/network/http/httptests"
	"github.com/Nivl/go-rest-tools/security/auth/testauth"
)

func TestToBasicAuth(t *testing.T) {
	_, sess := testauth.NewAuth()
	reqAuth := httptests.NewRequestAuth(sess)
	header := reqAuth.ToBasicAuth()

	authValue := fmt.Sprintf("%s:%s", sess.UserID, sess.ID)
	encoded := base64.StdEncoding.EncodeToString([]byte(authValue))
	expectedValue := "basic " + encoded
	assert.Equal(t, expectedValue, header, "invalid header returned")
}
