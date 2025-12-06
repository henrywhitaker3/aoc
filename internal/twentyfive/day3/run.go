// Package day3
package day3

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

type Bank []int

func (b Bank) LargestJoltage(digits int) int {
	out := []int{}
	candidates := slices.Clone(b)
	for len(out) < digits {
		slog.Debug("got candidates", "candidates", candidates, "out", out)
		dig, next := findNextDigit(candidates, digits-len(out))
		out = append(out, dig)
		candidates = next
		slog.Debug("picked out", "picked", dig)
	}

	arranged := ""
	for _, d := range out {
		arranged = fmt.Sprintf("%s%d", arranged, d)
	}
	dig, _ := strconv.Atoi(arranged)

	return dig
}

func findNextDigit(bank Bank, digits int) (int, Bank) {
	candidates := slices.Clone(bank)[0 : len(bank)-(digits-1)]
	index, digit := findLargestAtIndex(candidates)
	return digit, bank[index+1:]
}

func findLargestAtIndex(digits Bank) (int, int) {
	cloned := slices.Clone(digits)
	slices.Sort(cloned)
	slices.Reverse(cloned)
	slog.Debug("picking first digit", "from", cloned)
	return slices.Index(digits, cloned[0]), cloned[0]
}

func ParseData(data []byte) ([]Bank, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	out := []Bank{}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		line = strings.Trim(line, "\n")
		bank, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parse line: %w", err)
		}
		out = append(out, bank)
	}
	return out, nil
}

func parseLine(line string) (Bank, error) {
	out := Bank{}
	for _, char := range line {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, fmt.Errorf("parse char %s as int: %w", string(char), err)
		}
		out = append(out, val)
	}
	return out, nil
}

func SumLargestJoltages(banks []Bank, digits int) int {
	sums := []int{}
	for _, b := range banks {
		sums = append(sums, b.LargestJoltage(digits))
	}
	out := 0
	for _, s := range sums {
		out += s
	}
	return out
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	banks, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", SumLargestJoltages(banks, 2))

	return nil
}

func PartTwo(ctx context.Context) error {
	banks, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", SumLargestJoltages(banks, 12))

	return nil
}
