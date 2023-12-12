package pkg202310

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 4, A("example.txt"))
	require.Equal(t, 8, A("example2.txt"))
	require.Equal(t, 6968, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 1, B("example.txt"))
	require.Equal(t, 4, B("example3.txt"))
	require.Equal(t, 8, B("example4.txt"))
	require.Equal(t, 413, B("input.txt"))
}
