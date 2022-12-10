package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandom(t *testing.T) {
	s := RandomString(10)
	require.NotEmpty(t, s)
	email := RandomEmail()
	require.NotEmpty(t, email)
}
