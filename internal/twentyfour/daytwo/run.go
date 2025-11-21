// Package daytwo
package daytwo

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

type Report []int

func (r Report) Safe(tolerations int) bool {
	increasing := true

	for i, v := range r {
		if i == 0 {
			continue
		}

		diff := v - r[i-1]
		if i == 1 {
			if diff < 0 {
				increasing = false
			}
		}

		safe := r.isReportSafe(v, r[i-1], diff, increasing)
		if !safe {
			return false
		}
	}

	return true
}

func (r Report) isReportSafe(val int, prev int, diff int, increasing bool) bool {
	slog := slog.With("report", r, "val", val, "prev", prev, "diff", diff, "increasing", increasing)
	slog.Debug("checking report")
	absDiff := int(math.Abs(float64(diff)))
	if diff == 0 {
		return false
	}
	if diff > 0 && !increasing {
		slog.Debug("diff is positive, but first was decreasing")
		return false
	}
	if diff < 0 && increasing {
		slog.Debug("diff is negative, but first was increasing")
		return false
	}
	if absDiff < 1 {
		slog.Debug("skipping as diff less than 1")
		return false
	}
	if absDiff > 3 {
		slog.Debug("skipping as diff more than 3")
		return false
	}
	slog.Debug("got to the end")
	return true
}

type Reports []Report

func (r Reports) NumSafe(tolerations int) int {
	count := 0
	for _, r := range r {
		if r.Safe(tolerations) {
			count++
		}
	}
	return count
}

func loadData(input []byte) (Reports, error) {
	out := Reports{}
	r := bufio.NewReader(bytes.NewReader(input))
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read to new line : %w", err)
		}
		val, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parse line: %w", err)
		}
		out = append(out, val)
	}
	return out, nil
}

func parseLine(line string) (Report, error) {
	spl := strings.Split(strings.Trim(line, "\n"), " ")
	out := Report{}
	for _, v := range spl {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("parse int from line '%s': %w", line, err)
		}
		out = append(out, val)
	}
	return out, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	reports, err := loadData([]byte(input))
	if err != nil {
		return fmt.Errorf("load input: %w", err)
	}

	fmt.Printf("Total safe: %d\n", reports.NumSafe(0))

	return nil
}

func PartTwo(ctx context.Context) error {
	reports, err := loadData([]byte(input))
	if err != nil {
		return fmt.Errorf("load input: %w", err)
	}

	fmt.Printf("Total safe: %d\n", reports.NumSafe(1))

	return nil
}
