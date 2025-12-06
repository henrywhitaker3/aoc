// Package day5
package day5

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

type Database struct {
	ranges      [][]int
	ingredients []int
}

func (d Database) Fresh() []int {
	checks := []func(int) bool{}
	for _, r := range d.ranges {
		checks = append(checks, func(i int) bool {
			slog.Debug("checking against range", "start", r[0], "end", r[1], "i", i)
			return i >= r[0] && i <= r[1]
		})
	}

	out := []int{}
ol:
	for _, i := range d.ingredients {
		for _, check := range checks {
			if check(i) {
				out = append(out, i)
				continue ol
			}
		}
	}

	return out
}

func ParseData(data []byte) (*Database, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	out := &Database{}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		line = strings.Trim(line, "\n")
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			r, err := parseRange(line)
			if err != nil {
				return nil, fmt.Errorf("parse range: %w", err)
			}
			out.ranges = append(out.ranges, r)
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("parse id: %w", err)
			}
			out.ingredients = append(out.ingredients, id)
		}
	}
	return out, nil
}

func parseRange(line string) ([]int, error) {
	spl := strings.Split(line, "-")
	if len(spl) != 2 {
		return nil, fmt.Errorf("parse range should have 2 splits, had %d", len(spl))
	}
	start, err := strconv.Atoi(spl[0])
	if err != nil {
		return nil, fmt.Errorf("parse start '%s': %w", spl[0], err)
	}
	end, err := strconv.Atoi(spl[1])
	if err != nil {
		return nil, fmt.Errorf("parse end '%s': %w", spl[1], err)
	}
	return []int{start, end}, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	db, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Count: %d\n", len(db.Fresh()))

	return nil
}

func PartTwo(ctx context.Context) error {
	return nil
}
