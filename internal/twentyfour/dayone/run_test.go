package dayone

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed sample.txt
	sample string
)

func TestItParsesSample(t *testing.T) {
	entries, err := loadData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 11, entries.Total())
}

func TestItParsesSampleSimilarity(t *testing.T) {
	entries, err := loadData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 31, entries.Similarity())
}
