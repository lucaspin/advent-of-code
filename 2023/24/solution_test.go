package pkg202324

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 2, A("example.txt", 7.0, 27.0))
	require.Equal(t, 2, A("input.txt", 200000000000000.0, 400000000000000.0))
}
