package pkg202307

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, int64(6440), A("test.txt"))
	require.Equal(t, int64(250347426), A("a.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, int64(5905), B("test.txt"))
	require.Equal(t, int64(251224870), B("a.txt"))
}
