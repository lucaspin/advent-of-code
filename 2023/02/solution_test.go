package pkg202302

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 8, A(Bag{Red: 12, Green: 13, Blue: 14}, "test.txt"))
	require.Equal(t, 2632, A(Bag{Red: 12, Green: 13, Blue: 14}, "a.txt"))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, int64(2286), B(Bag{Red: 12, Green: 13, Blue: 14}, "test.txt"))
	require.Equal(t, int64(2286), B(Bag{Red: 12, Green: 13, Blue: 14}, "a.txt"))
}
