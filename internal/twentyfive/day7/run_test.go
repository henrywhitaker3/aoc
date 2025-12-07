package day7

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
	man, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 240, len(man))
	require.Equal(t, 7, man.Start().X)
}

func TestItCountsSplits(t *testing.T) {
	common.TestLogger(t)
	man, err := ParseData([]byte(sample))
	require.Nil(t, err)

	require.Equal(t, 21, CountSplits(man))
}
