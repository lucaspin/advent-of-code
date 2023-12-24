package pkg202323

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 94, A("example.txt"))
	require.Equal(t, 2206, A("input.txt"))
}

func Test__SolutionTryB(t *testing.T) {
	require.Equal(t, 154, B("example.txt"))
	require.Equal(t, 6490, B("input.txt"))
}
