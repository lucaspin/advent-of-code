package pkg202312

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 21, A("example.txt"))
	require.Equal(t, 7633, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 525152, B("example.txt"))
	require.Equal(t, 23903579139437, B("input.txt"))
}
