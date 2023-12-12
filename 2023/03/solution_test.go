package pkg202303

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__Solution(t *testing.T) {
	require.Equal(t, int64(4361), A("test.txt"))
	require.Equal(t, int64(537832), A("a.txt"))
	require.Equal(t, int64(467835), B("test.txt"))
	require.Equal(t, int64(81939900), B("a.txt"))
}
