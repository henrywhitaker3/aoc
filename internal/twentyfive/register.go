// Package twentyfive
package twentyfive

import (
	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/henrywhitaker3/aoc/internal/twentyfive/dayone"
)

func Register(r common.Registerer) {
	r.Set(2025, 1, 1, dayone.PartOne)
	r.Set(2025, 1, 2, dayone.PartTwo)
}
