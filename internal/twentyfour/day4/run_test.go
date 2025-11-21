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
	common.TestLogger(t)
	graph, err := ParseData([]byte(sample))
	require.Nil(t, err)
	t.Log(graph)
}

func TestItFindsXmases(t *testing.T) {
	common.TestLogger(t)
	graph, err := ParseData([]byte(sample))
	require.Nil(t, err)
	count, err := CountXmas(graph)
	require.Nil(t, err)
	require.Equal(t, 18, count)
}

func TestItFindsCrossMas(t *testing.T) {
	common.TestLogger(t)
	graph, err := ParseData([]byte(sample))
	require.Nil(t, err)
	count, err := CountCrossMas(graph)
	require.Nil(t, err)
	require.Equal(t, 9, count)
}
