package pkg202317

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 102, A("example.txt"))
	require.Equal(t, 102, A("input.txt"))
}
