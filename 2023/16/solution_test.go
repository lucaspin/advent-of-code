package pkg202316

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 46, A("example.txt"))
	require.Equal(t, 7939, A("input.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 51, B("example.txt"))
	require.Equal(t, 8318, B("input.txt"))
}
