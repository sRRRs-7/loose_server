package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandom(t *testing.T) {
	s := RandomString(10)
	require.NotEmpty(t, s)

	email := RandomEmail()
	require.Contains(t, email, "@email.com")
	require.NotEmpty(t, email)

	n := RandomInteger(1, 10)
	require.True(t, n <= 10)
	require.NotEmpty(t, n)
}
