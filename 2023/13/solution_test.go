package pkg202313

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 405, A("example.txt"))
	require.Equal(t, 37718, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	// require.Equal(t, 400, B("example.txt"))
	// require.Equal(t, 37718, A("input.txt"))
}
