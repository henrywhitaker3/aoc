package daytwo

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

func TestItCountsSafeReports(t *testing.T) {
	common.TestLogger(t)
	reports, err := loadData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 2, reports.NumSafe(0))
}

func TestItCountsSafeReportsWith1Toleration(t *testing.T) {
	common.TestLogger(t)
	reports, err := loadData([]byte(sample))
	require.Nil(t, err)
	require.Equal(t, 4, reports.NumSafe(1))
}
