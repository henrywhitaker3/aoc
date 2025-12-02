// Package day2
package day2

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type IDRange []int

func ParseData(data []byte) ([]IDRange, error) {
	lines := strings.Split(strings.ReplaceAll(string(data), "\n", ""), ",")
	out := []IDRange{}
	for _, line := range lines {
		spl := strings.Split(line, "-")
		start, err := strconv.Atoi(spl[0])
		if err != nil {
			return nil, fmt.Errorf("parse start: %w", err)
		}
		end, err := strconv.Atoi(spl[1])
		if err != nil {
			return nil, fmt.Errorf("parse end: %w", err)
		}
		out = append(out, IDRange{start, end})
	}
	return out, nil
}

func ValidateRanges(r []IDRange, extended bool) ([]int, []int) {
	valid := []int{}
	invalid := []int{}

	var validator func(int) bool
	switch extended {
	case true:
		validator = isValidIDExtended
	case false:
		validator = isValidID
	}

	for _, r := range r {
		for i := r[0]; i <= r[1]; i++ {
			if validator(i) {
				valid = append(valid, i)
			} else {
				invalid = append(invalid, i)
			}
		}
	}

	return valid, invalid
}

func SumInvalidIDs(r []IDRange, extended bool) int {
	_, invalid := ValidateRanges(r, extended)
	out := 0
	for _, i := range invalid {
		slog.Debug("found invalid id", "id", i)
		out += i
	}
	return out
}

func isValidID(id int) bool {
	// All ids with an odd number of digits are valid
	if countDigits(id)%2 != 0 {
		return true
	}

	asStr := fmt.Sprintf("%d", id)
	first := asStr[0:(len(asStr) / 2)]
	second := asStr[(len(asStr) / 2):]
	slog.Debug("checking halves", "raw", asStr, "first", first, "second", second)

	return first != second
}

func isValidIDExtended(id int) bool {
	if !isValidID(id) {
		return false
	}

	asStr := fmt.Sprintf("%d", id)

	for i := range countDigits(id) / 2 {
		match := asStr[:i+1]
		slog.Debug("checking match", "id", id, "match", match)
		if strings.ReplaceAll(asStr, match, "") == "" {
			return false
		}
	}

	return true
}

func countDigits(id int) int {
	count := 0
	for id > 0 {
		id = id / 10
		count++
	}
	return count
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	ranges, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", SumInvalidIDs(ranges, false))

	return nil
}

func PartTwo(ctx context.Context) error {
	ranges, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", SumInvalidIDs(ranges, true))

	return nil
}
