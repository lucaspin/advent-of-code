package pkg202321

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, int64(3598), A("input.txt", 64))
}

func Test__SolutionB(t *testing.T) {
	require.Equal(t, int64(601441063166538), B("input.txt", 26501365))
}
