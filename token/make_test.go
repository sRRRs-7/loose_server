package token

import (
	"testing"
	"time"

	"github.com/sRRRs-7/loose_style.git/utils"
	"github.com/stretchr/testify/require"
)

func TestMaker(t *testing.T) {
	payload, err := NewPayload(utils.RandomString(10), 1*time.Second)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	err = payload.ValidToken()
	require.NoError(t, err)

	b, err := payload.MarshalBinary()
	require.NoError(t, err)
	require.NotEmpty(t, b)

	err = payload.UnmarshalBinary(b)
	require.NoError(t, err)
}
