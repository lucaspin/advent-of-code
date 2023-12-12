package pkg202301

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test__First(t *testing.T) {
	require.Equal(t, int64(142), First([]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}))
	f, err := os.ReadFile("./input.txt")
	require.NoError(t, err)
	lines := strings.Split(string(f), "\n")
	require.Equal(t, int64(54940), First(lines))
}

func Test__Second(t *testing.T) {
	require.Equal(t, int64(281), Second([]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}))
	f, err := os.ReadFile("./input.txt")
	require.NoError(t, err)
	lines := strings.Split(string(f), "\n")
	require.Equal(t, int64(54265), Second(lines))
}
