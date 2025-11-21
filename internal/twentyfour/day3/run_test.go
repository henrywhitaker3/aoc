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
	//go:embed sample2.txt
	sample2 string
)

func TestItMultpipliesSamplePartOne(t *testing.T) {
	common.TestLogger(t)
	val, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 161, val)
}

func TestItMultpipliesSamplePartTwo(t *testing.T) {
	common.TestLogger(t)
	val, err := ParseDataWithSwitch([]byte(sample2))
	require.Nil(t, err)
	require.Equal(t, 48, val)
}
