package day1

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
	turns, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 10, len(turns))
	require.Equal(t, Direction("L"), turns[0].Direction)
	require.Equal(t, 68, turns[0].Quantity)
}

func TestItGetsExpectedTurns(t *testing.T) {
	common.TestLogger(t)
	turns, err := ParseData([]byte(sample))
	require.Nil(t, err)
	expected := []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
	i := 0
	dial := NewDial(50)
	dial.Go(turns, func(pos int) {
		require.Equal(t, expected[i], pos)
		i++
	})
}

func TestItCountsZeros(t *testing.T) {
	turns, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 3, CountZeros(turns, 50))
}
