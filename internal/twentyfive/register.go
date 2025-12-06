// Package twentyfive
package twentyfive

import (
	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/henrywhitaker3/aoc/internal/twentyfive/day1"
	"github.com/henrywhitaker3/aoc/internal/twentyfive/day2"
	"github.com/henrywhitaker3/aoc/internal/twentyfive/day3"
	"github.com/henrywhitaker3/aoc/internal/twentyfive/day4"
)

func Register(r common.Registerer) {
	r.Set(2025, 1, 1, day1.PartOne)
	r.Set(2025, 1, 2, day1.PartTwo)
	r.Set(2025, 2, 1, day2.PartOne)
	r.Set(2025, 2, 2, day2.PartTwo)
	r.Set(2025, 3, 1, day3.PartOne)
	r.Set(2025, 3, 2, day3.PartTwo)
	r.Set(2025, 4, 1, day4.PartOne)
	r.Set(2025, 4, 2, day4.PartTwo)
}
