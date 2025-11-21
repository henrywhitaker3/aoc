// Package twentyfour
package twentyfour

import (
	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day1"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day2"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day3"
)

func Register(r common.Registerer) {
	r.Set(2024, 1, 1, day1.PartOne)
	r.Set(2024, 1, 2, day1.PartTwo)
	r.Set(2024, 2, 1, day2.PartOne)
	r.Set(2024, 2, 2, day2.PartTwo)
	r.Set(2024, 3, 1, day3.PartOne)
	r.Set(2024, 3, 2, day3.PartTwo)
}
