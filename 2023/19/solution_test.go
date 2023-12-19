package pkg202319

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__SolutionA(t *testing.T) {
	require.Equal(t, 19114, A("example.txt"))
	require.Equal(t, 319295, A("input.txt"))
}
