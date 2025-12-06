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
		[]Operand[int]{
			{
				parsed: 123,
			},
			{
				parsed: 45,
			},
			{
				parsed: 6,
			},
		},
		calcs[0].Numbers,
	)
	require.Equal(t, Operand[Operator]{parsed: Multiply}, calcs[0].Operator)
}

func TestItSumsAnswers(t *testing.T) {
	calcs, err := ParseData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 4277556, SumResults(calcs))
}
