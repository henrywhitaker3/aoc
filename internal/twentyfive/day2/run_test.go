package day2

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed sample.txt
	sample string
)

func TestItParsesData(t *testing.T) {
	ranges, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 11, len(ranges))
	require.Equal(t, 2121212124, ranges[10][1])
}

func TestItSumsInvalidIDs(t *testing.T) {
	common.TestLogger(t)
	ranges, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 1227775554, SumInvalidIDs(ranges))
}

func TestItCountsDigitsInNumbers(t *testing.T) {
	tcs := []struct {
		input    int
		expected int
	}{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    12,
			expected: 2,
		},
		{
			input:    123,
			expected: 3,
		},
		{
			input:    1234567,
			expected: 7,
		},
	}

	for _, c := range tcs {
		t.Run(fmt.Sprintf("%d", c.input), func(t *testing.T) {
			require.Equal(t, c.expected, countDigits(c.input))
		})
	}
}
