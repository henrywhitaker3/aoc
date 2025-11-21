// Package twentyfive
package twentyfive

import (
	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/henrywhitaker3/aoc/internal/twentyfive/day1"
)

func Register(r common.Registerer) {
	r.Set(2025, 1, 1, day1.PartOne)
	r.Set(2025, 1, 2, day1.PartTwo)
}
