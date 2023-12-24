package pkg202324

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 2, A("example.txt", 7.0, 27.0))
	require.Equal(t, 31921, A("input.txt", 200000000000000.0, 400000000000000.0))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, int64(47), B("example.txt"))
	require.Equal(t, int64(761691907059631), B("input.txt"))
}
