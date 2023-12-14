package pkg202314

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 136, A("example.txt"))
	require.Equal(t, 109654, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 64, B("example.txt"))
	require.Equal(t, 94876, B("input.txt"))
}
