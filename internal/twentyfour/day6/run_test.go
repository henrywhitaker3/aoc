package day6

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

func TestItParsesInput(t *testing.T) {
	m, err := ParseData([]byte(sample))
	require.Nil(t, err)

	require.Equal(t, 4, m.guard.X)
	require.Equal(t, 6, m.guard.Y)
}

func TestItCalculatesPartOneMoves(t *testing.T) {
	common.TestLogger(t)
	m, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 41, m.SumMoves())
}

func TestItChangesDirection(t *testing.T) {
	// Move up
	start := []int{0, -1}

	t.Run("up changes to right", func(t *testing.T) {
		changed := changeDirection(start)
		require.Equal(t, []int{1, 0}, changed)
	})

	t.Run("right changes to down", func(t *testing.T) {
		changed := changeDirection([]int{1, 0})
		require.Equal(t, []int{0, 1}, changed)
	})

	t.Run("down changes to left", func(t *testing.T) {
		changed := changeDirection([]int{0, 1})
		require.Equal(t, []int{-1, 0}, changed)
	})

	t.Run("left changes to up", func(t *testing.T) {
		changed := changeDirection([]int{-1, 0})
		require.Equal(t, []int{0, -1}, changed)
	})
}
