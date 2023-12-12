package pkg202311

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 374, Solution("example.txt", 2))
	require.Equal(t, 9947476, Solution("input.txt", 2))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 374, Solution("example.txt", 2))
	require.Equal(t, 1030, Solution("example.txt", 10))
	require.Equal(t, 8410, Solution("example.txt", 100))
	require.Equal(t, 519939907614, Solution("input.txt", 1000000))
}
