package day7

import (
	"context"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed sample.txt
	sample string
)

func TestItParsesData(t *testing.T) {
	eqs, err := ParseData([]byte(sample))
	require.Nil(t, err)

	require.Equal(t, 190, eqs[0].Answer)
	require.Equal(t, 10, eqs[0].Inputs[0])
	require.Equal(t, 19, eqs[0].Inputs[1])
}

func TestItSumsEquations(t *testing.T) {
	eqs, err := ParseData([]byte(sample))
	require.Nil(t, err)

	require.Equal(t, 3749, Evaluate(context.Background(), eqs, false))
}

func TestItEvaluatesEquations(t *testing.T) {
	tcs := []struct {
		name string
		eq   Equation
	}{
		{
			name: "evaluates two inputs",
			eq: Equation{
				Answer: 190,
				Inputs: []int{10, 19},
			},
		},
		{
			name: "evaluates three inputs",
			eq: Equation{
				Answer: 3267,
				Inputs: []int{81, 40, 27},
			},
		},
		{
			name: "evaluates four inputs",
			eq: Equation{
				Answer: 292,
				Inputs: []int{11, 6, 16, 20},
			},
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			require.True(t, evaluate(c.eq.Answer, c.eq.Inputs, 0, false))
		})
	}
}
