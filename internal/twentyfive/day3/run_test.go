package day3

import (
	_ "embed"
	"testing"

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
	banks, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 98, banks[0].LargestJoltage())
	require.Equal(t, 89, banks[1].LargestJoltage())
	require.Equal(t, 78, banks[2].LargestJoltage())
	require.Equal(t, 92, banks[3].LargestJoltage())
}

func TestItSumsLargestJoltages(t *testing.T) {
	banks, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 357, SumLargestJoltages(banks))
}
