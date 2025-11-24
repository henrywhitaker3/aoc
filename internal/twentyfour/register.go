// Package twentyfour
package twentyfour

import (
	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day1"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day2"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day3"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day4"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day5"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day6"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/day7"
)

func Register(r common.Registerer) {
	r.Set(2024, 1, 1, day1.PartOne)
	r.Set(2024, 1, 2, day1.PartTwo)
	r.Set(2024, 2, 1, day2.PartOne)
	r.Set(2024, 2, 2, day2.PartTwo)
	r.Set(2024, 3, 1, day3.PartOne)
	r.Set(2024, 3, 2, day3.PartTwo)
	r.Set(2024, 4, 1, day4.PartOne)
	r.Set(2024, 4, 2, day4.PartTwo)
	r.Set(2024, 5, 1, day5.PartOne)
	r.Set(2024, 5, 2, day5.PartTwo)
	r.Set(2024, 6, 1, day6.PartOne)
	r.Set(2024, 6, 2, day6.PartTwo)
	r.Set(2024, 7, 1, day7.PartOne)
	r.Set(2024, 7, 2, day7.PartTwo)
}
