package bcrypt_test

import (
	"testing"

	"github.com/Nivl/go-rest-tools/security/hasher/bcrypt"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	hasher := &bcrypt.Bcrypt{}
	psw := "\\My secuRe p@5word."

	cryptedPsw, err := hasher.Hash(psw)
	assert.NoError(t, err, "CryptPassword() should have worked")

	// test a valid password
	isValid := hasher.IsValid(cryptedPsw, psw)
	assert.True(t, isValid, "IsPasswordValid() should have returned true")

	// test an invalid password
	isValid = hasher.IsValid(cryptedPsw, "invalid password")
	assert.False(t, isValid, "IsPasswordValid() should have returned false")

}
