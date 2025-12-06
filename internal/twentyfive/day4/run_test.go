package day4

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
	grid, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 100, len(grid))
	point, ok := grid.Find(4, 5)
	require.True(t, ok)
	require.True(t, point.Paper)
}

func TestItCountsAdjacentPaper(t *testing.T) {
	common.TestLogger(t)
	grid, err := ParseData([]byte(sample))
	require.Nil(t, err)

	zerozero, ok := grid.Find(0, 0)
	require.True(t, ok)
	require.Equal(t, 2, grid.AdjacentPaper(zerozero))
	onezero, ok := grid.Find(1, 0)
	require.True(t, ok)
	require.Equal(t, 4, grid.AdjacentPaper(onezero))
	fourfour, ok := grid.Find(4, 4)
	t.Log(grid.adjacentPoints(fourfour))
	require.True(t, ok)
	require.Equal(t, 8, grid.AdjacentPaper(fourfour))
}

func TestItFindsMovablePoints(t *testing.T) {
	common.TestLogger(t)
	grid, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 13, len(grid.MoveablePoints()))
}
