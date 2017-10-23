package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Nivl/go-rest-tools/security/auth"
)

func TestUserIsLogged(t *testing.T) {
	var u *auth.User
	assert.False(t, u.IsLogged(), "IsLogged() should have returned false")

	u = &auth.User{}
	assert.True(t, u.IsLogged(), "IsLogged() should have returned true")
}

func TestUserIsAdm(t *testing.T) {
	var u *auth.User
	assert.False(t, u.IsAdm(), "IsLogged() should have returned false")

	u = &auth.User{}
	assert.False(t, u.IsAdm(), "IsLogged() should have returned false")

	u = &auth.User{IsAdmin: true}
	assert.True(t, u.IsAdm(), "IsLogged() should have returned true")
}
