// Package dayone
package dayone

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Entry struct {
	Left  int
	Right int
}

func (e Entry) Distance() int {
	return int(math.Abs(float64(e.Left) - float64(e.Right)))
}

type Entries []Entry

func (e Entries) Total() int {
	out := 0
	for _, e := range e {
		out += e.Distance()
	}
	return out
}

func (e Entries) Similarity() int {
	heat := map[int]int{}
	for _, v := range e {
		heat[v.Right]++
	}

	out := 0

	for _, v := range e {
		count := 0
		if val, ok := heat[v.Left]; ok {
			count = val
		}
		out += (v.Left * count)
	}

	return out
}

func loadData(input []byte) (Entries, error) {
	left := []int{}
	right := []int{}

	r := bufio.NewReader(bytes.NewReader(input))

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}

		leftInt, rightInt, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parse line '%s': %w", line, err)
		}
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	slices.Sort(left)
	slices.Sort(right)

	if len(left) != len(right) {
		return nil, fmt.Errorf("sizes of slices are different")
	}

	out := Entries{}
	for i := range left {
		out = append(out, Entry{
			Left:  left[i],
			Right: right[i],
		})
	}

	return out, nil
}

func parseLine(line string) (int, int, error) {
	line = strings.Trim(line, "\n")
	spl := strings.Split(line, " ")
	left, err := strconv.Atoi(spl[0])
	if err != nil {
		return 0, 0, fmt.Errorf("parse left int: %w", err)
	}
	right, err := strconv.Atoi(spl[len(spl)-1])
	if err != nil {
		return 0, 0, fmt.Errorf("parse right int: %w", err)
	}
	return left, right, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	entries, err := loadData([]byte(input))
	if err != nil {
		return fmt.Errorf("load input.txt: %w", err)
	}

	fmt.Printf("Total: %d\n", entries.Total())

	return nil
}

func PartTwo(ctx context.Context) error {
	entries, err := loadData([]byte(input))
	if err != nil {
		return fmt.Errorf("load input.txt: %w", err)
	}

	fmt.Printf("Similarity: %d\n", entries.Similarity())

	return nil
}
