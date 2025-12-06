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

func TestItParsesData(t *testing.T) {
	common.TestLogger(t)
	calcs, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(
		t,
		[]Operand{
			123,
			45,
			6,
		},
		calcs[0].Numbers,
	)
	require.Equal(t, Multiply, calcs[0].Operator)
}

func TestItSumsAnswers(t *testing.T) {
	calcs, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 4277556, SumResults(calcs))
}

func TestItSumsRTLAnswers(t *testing.T) {
	calcs, err := ParseRTL([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 3263827, SumResults(calcs))
}

func TestItParsesRTLData(t *testing.T) {
	common.TestLogger(t)
	calcs, err := ParseRTL([]byte(sample))
	require.Nil(t, err)
	require.Equal(
		t,
		[]Operand{
			4,
			431,
			623,
		},
		calcs[0].Numbers,
	)
	require.Equal(t, Add, calcs[0].Operator)
}
