package cfg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	config, err := LoadConfig(".")
	fmt.Println(config
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
