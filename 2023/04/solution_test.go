package pkg202304

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__Solution(t *testing.T) {
	require.Equal(t, int64(13), A("test.txt"))
	require.Equal(t, int64(21485), A("a.txt"))
	require.Equal(t, int64(30), B("test.txt"))
	require.Equal(t, int64(11024379), B("a.txt"))
}
