package pkg202323

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 94, A("example.txt"))
	require.Equal(t, 2206, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 154, B("example.txt"))
	require.Equal(t, 0, B("input.txt"))
}
