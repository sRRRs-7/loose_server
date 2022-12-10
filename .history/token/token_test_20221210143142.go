package token

import (
	"testing"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	m, err := NewPasetoMaker(utils.RandomString(10))
	require.NoError(t, err)
	require.NotEmpty(t, m)
}
