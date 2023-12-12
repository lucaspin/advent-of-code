package pkg202306

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 288, A("test.txt"))
	require.Equal(t, 512295, A("a.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, 71503, B("test.txt"))
	require.Equal(t, 36530883, B("a.txt"))
}
