// Package twentyfour
package twentyfour

import (
	"github.com/henrywhitaker3/aoc/internal/common"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/dayone"
	"github.com/henrywhitaker3/aoc/internal/twentyfour/daytwo"
)

func Register(r common.Registerer) {
	r.Set(2024, 1, 1, dayone.PartOne)
	r.Set(2024, 1, 2, dayone.PartTwo)
	r.Set(2024, 2, 1, daytwo.PartOne)
	r.Set(2024, 2, 2, daytwo.PartTwo)
}
