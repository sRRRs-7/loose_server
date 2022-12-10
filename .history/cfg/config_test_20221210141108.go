package cfg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	config, err := LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
