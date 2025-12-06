// Package day3
package day3

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Bank []int

func (b Bank) LargestJoltage() int {
	cloned := slices.Clone(b)
	slices.Sort(cloned)
	slices.Reverse(cloned)

	pick := 0
	for {
		val := cloned[pick]
		if slices.Index(b, val) == (len(b) - 1) {
			// Move to the next one, there's no other number to put next if
			// it is at the end
			pick++
			continue
		}
		break
	}

	first := cloned[pick]
	remaining := slices.Clone(b[slices.Index(b, first)+1:])
	slices.Sort(remaining)
	slices.Reverse(remaining)
	second := b[slices.Index(b, remaining[0])]

	out, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, second))
	return out
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

func SumLargestJoltages(banks []Bank) int {
	sums := []int{}
	for _, b := range banks {
		sums = append(sums, b.LargestJoltage())
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

	fmt.Printf("Sum: %d\n", SumLargestJoltages(banks))

	return nil
}

func PartTwo(ctx context.Context) error {
	return nil
}
