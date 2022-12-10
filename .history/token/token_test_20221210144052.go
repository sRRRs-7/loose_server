package token

import (
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	m, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, m)

	token, payload, err := m.CreateToken(utils.RandomString(10), 1*time.Second)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.NotEmpty(t, s)
	require.Contains(t, s, "v2.local")

	m.VerifyToken()
}
