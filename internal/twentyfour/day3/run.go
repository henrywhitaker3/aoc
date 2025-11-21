// Package day3
package day3

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

var (
	mulRegex   = regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))`)
	doMulRegex = regexp.MustCompile(`(do\(\)|don't\(\))|(mul\([0-9]{1,3},[0-9]{1,3}\))`)
)

func ParseData(data []byte) (int, error) {
	matches := mulRegex.FindAllString(string(data), -1)
	slog.Debug("got matches", "matches", matches)

	out := 0

	for _, m := range matches {
		val, err := processMul(m)
		if err != nil {
			return 0, fmt.Errorf("process match: %w", err)
		}
		out += val
	}

	return out, nil
}

func ParseDataWithSwitch(data []byte) (int, error) {
	matches := doMulRegex.FindAllString(string(data), -1)
	slog.Debug("got matches", "matches", matches)

	out := 0
	do := true

	for _, m := range matches {
		switch m {
		case doStr:
			do = true
			continue
		case dont:
			do = false
			continue
		}
		if do {
			val, err := processMul(m)
			if err != nil {
				return 0, fmt.Errorf("process match: %w", err)
			}
			out += val
		}
	}

	return out, nil
}

const (
	doStr = "do()"
	dont  = "don't()"
)

func processMul(match string) (int, error) {
	match = strings.Replace(match, "mul(", "", 1)
	match = strings.Replace(match, ")", "", 1)
	spl := strings.Split(match, ",")
	if len(spl) != 2 {
		return 0, fmt.Errorf("got invalid number in split for line '%s'", match)
	}

	left, err := strconv.Atoi(spl[0])
	if err != nil {
		return 0, fmt.Errorf("could not parse left int: %w", err)
	}
	right, err := strconv.Atoi(spl[1])
	if err != nil {
		return 0, fmt.Errorf("could not parse right int: %w", err)
	}

	return left * right, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	sum, err := ParseData([]byte(input))
	if err != nil {
		return err
	}

	fmt.Printf("Sum: %d\n", sum)

	return nil
}

func PartTwo(ctx context.Context) error {
	sum, err := ParseDataWithSwitch([]byte(input))
	if err != nil {
		return err
	}

	fmt.Printf("Sum: %d\n", sum)

	return nil
}
