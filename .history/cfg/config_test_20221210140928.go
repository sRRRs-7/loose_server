package cfg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	config, err := LoadConfig(".")
	fmt.println(err)
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
