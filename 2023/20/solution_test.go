package pkg202320

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, int64(32000000), A("example.txt"))
	require.Equal(t, int64(11687500), A("example2.txt"))
	require.Equal(t, int64(743090292), A("input.txt"))
}
