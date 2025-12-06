package day3

import (
	_ "embed"
	"testing"

	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed sample.txt
	sample string
)

func TestItParsesData(t *testing.T) {
	banks, err := ParseData([]byte(sample))
	require.Nil(t, err)
	t.Log(banks)
	require.Equal(t, 4, len(banks))
	require.Equal(t, 9, banks[0][0])
}

func TestItFindsLargestJoltages(t *testing.T) {
	common.TestLogger(t)
	banks, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 98, banks[0].LargestJoltage(2))
	require.Equal(t, 89, banks[1].LargestJoltage(2))
	require.Equal(t, 78, banks[2].LargestJoltage(2))
	require.Equal(t, 92, banks[3].LargestJoltage(2))
}

func TestItSums2LargestJoltages(t *testing.T) {
	banks, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 357, SumLargestJoltages(banks, 2))
}

func TestItSums12LargestJoltages(t *testing.T) {
	banks, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 3121910778619, SumLargestJoltages(banks, 12))
}
