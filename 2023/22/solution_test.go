package pkg202322

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 5, A("example.txt"))
	require.Equal(t, 395, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 7, B("example.txt"))
	require.Equal(t, 64714, B("input.txt"))
}
