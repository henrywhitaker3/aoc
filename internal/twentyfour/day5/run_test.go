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
	common.TestLogger(t)
	rules, updates, err := ParseData([]byte(sample))
	require.Nil(t, err)

	require.Equal(t, 21, len(rules))
	require.Equal(t, 47, rules[0].Left)
	require.Equal(t, 53, rules[len(rules)-1].Left)

	require.Equal(t, 6, len(updates))
	require.Equal(t, 75, updates[0][0])
	require.Equal(t, 47, updates[len(updates)-1][4])
}

func TestItGetsTheIncorrectUpdates(t *testing.T) {
	common.TestLogger(t)
	rules, updates, err := ParseData([]byte(sample))
	require.Nil(t, err)

	incorrect := rules.GetCorrectedUpdates(updates)
	require.Equal(t, 3, len(incorrect))
	require.Equal(t, 123, sumMiddleValues(incorrect))
}

func TestItGetsTheUpdates(t *testing.T) {
	common.TestLogger(t)
	rules, updates, err := ParseData([]byte(sample))
	require.Nil(t, err)

	correct := rules.GetCorrectUpdates(updates)
	require.Equal(t, 3, len(correct))
	require.Equal(t, 143, sumMiddleValues(correct))
}

func TestItAssessRulesCorrectly(t *testing.T) {
	tcs := []struct {
		name   string
		rule   Rule
		update Update
		passes bool
	}{
		{
			name: "passes when both pages are not in update",
			rule: Rule{
				Left:  53,
				Right: 84,
			},
			update: Update{7, 9, 53, 80},
			passes: true,
		},
		{
			name: "passes when left is in update first",
			rule: Rule{
				Left:  53,
				Right: 84,
			},
			update: Update{7, 9, 53, 80, 84},
			passes: true,
		},
		{
			name: "does not pass when left is in update first",
			rule: Rule{
				Left:  84,
				Right: 53,
			},
			update: Update{7, 9, 53, 80, 84},
			passes: false,
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.passes, c.rule.Passes(c.update))
		})
	}
}
