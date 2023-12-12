package pkg202305

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 35, A("test.txt"))
	require.Equal(t, 177942185, A("a.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 46, B("test.txt"))
	require.Equal(t, 69841803, B("a.txt"))
}
