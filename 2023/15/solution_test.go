package pkg202315

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 1320, A("example.txt"))
	require.Equal(t, 1, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 145, B("example.txt"))
	require.Equal(t, 145, B("input.txt"))
}
