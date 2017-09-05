package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Nivl/go-rest-tools/security/auth"
)

func TestPassword(t *testing.T) {
	psw := "\\My secuRe p@5word."

	cryptedPsw, err := auth.CryptPassword(psw)
	assert.NoError(t, err, "CryptPassword() should have worked")

	// test a valid password
	isValid := auth.IsPasswordValid(cryptedPsw, psw)
	assert.True(t, isValid, "IsPasswordValid() should have returned true")

	// test an invalid password
	isValid = auth.IsPasswordValid(cryptedPsw, "invalid password")
	assert.False(t, isValid, "IsPasswordValid() should have returned false")

}

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
