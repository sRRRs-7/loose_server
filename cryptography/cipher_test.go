package cryptography

import (
	"testing"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)


func TestPasswords(t *testing.T) {
	// input password
	password := utils.RandomString(5)

	// hash password
	hashPassword1, err := HashPassword(password)
	require.NoError(t,err)
	require.NotEmpty(t,hashPassword1)

	hashPassword2, err := HashPassword(password)
	require.NoError(t,err)
	require.NotEmpty(t,hashPassword2)

	b, err := VerifyHash(hashPassword1, hashPassword2)
	require.NoError(t, err)
	require.Equal(t, b, true)

	// wrong password
	wrongPassword, err := HashPassword(password + "salt")
	require.NoError(t, err)
	require.NotEqual(t, hashPassword1, wrongPassword)

	b, err = VerifyHash(hashPassword1, wrongPassword)
	require.Error(t, err)
	require.NotEqual(t, b, true)
}

