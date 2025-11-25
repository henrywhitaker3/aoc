package day8

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
	m, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, false, m.points[0].Broadcasting())
	require.Equal(t, []string{"0", "A"}, m.uniqueFrequencies)
}

func TestItFindsAntinodes(t *testing.T) {
	common.TestLogger(t)
	m, err := ParseData([]byte(sample))
	require.Nil(t, err)

	allAntinodes := []Point{}
	for _, freq := range m.uniqueFrequencies {
		nodes := m.Antinodes(freq)
		allAntinodes = append(allAntinodes, nodes...)
	}

	filtered := Unique(allAntinodes)
	t.Log(filtered)
	require.Equal(t, 14, len(filtered))
}
