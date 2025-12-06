package day5

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
	db, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 4, len(db.ranges))
	require.Equal(t, 6, len(db.ingredients))
	require.Equal(
		t,
		[][]int{
			{3, 5},
			{10, 14},
			{16, 20},
			{12, 18},
		},
		db.ranges,
	)
}

func TestItCollectsFreshIngredients(t *testing.T) {
	common.TestLogger(t)
	db, err := ParseData([]byte(sample))
	require.Nil(t, err)

	require.Equal(t, 3, len(db.Fresh()))
}
