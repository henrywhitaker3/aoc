// Package day8
package day8

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"math"
	"slices"
	"strings"
)

type Point struct {
	X         int
	Y         int
	Frequency string
}

func (p Point) Broadcasting() bool {
	return p.Frequency != ""
}

type Map struct {
	points            []Point
	uniqueFrequencies []string
}

func (m *Map) Antinodes(freq string, attempts int) []Point {
	out := []Point{}

	for i, a := range m.points {
		if a.Frequency != freq {
			continue
		}
		for j, b := range m.points {
			if i == j || b.Frequency != freq {
				continue
			}

			if attempts == int(math.Inf(1)) {
				out = append(out, a)
				out = append(out, b)
			}

			adX := a.X - b.X
			adY := a.Y - b.Y
			bdX := adX * -1
			bdY := adY * -1

			slog := slog.With(
				"a",
				fmt.Sprintf("%d,%d", a.X, a.Y),
				"b",
				fmt.Sprintf("%d,%d", b.X, b.Y),
			)

			for k := range attempts {
				slog.Debug("checking", "attempt", k+1)
				adX := adX * (k + 1)
				adY := adY * (k + 1)
				bdX := bdX * (k + 1)
				bdY := bdY * (k + 1)

				found := false
				if p, ok := m.Find(a.X+adX, a.Y+adY); ok {
					slog.Debug("found antinode", "node", fmt.Sprintf("%d,%d", p.X, p.Y))
					out = append(out, p)
					found = true
				}
				if p, ok := m.Find(b.X+bdX, b.Y+bdY); ok {
					slog.Debug("found antinode", "node", fmt.Sprintf("%d,%d", p.X, p.Y))
					out = append(out, p)
					found = true
				}
				if !found {
					slog.Debug("breaking")
					break
				}
			}
		}
	}

	return Unique(out)
}

func Unique(input []Point) []Point {
	out := []Point{}
	for _, p := range input {
		if !slices.ContainsFunc(out, func(t Point) bool {
			return p.X == t.X && p.Y == t.Y
		}) {
			out = append(out, p)
		}
	}
	return out
}

func (m *Map) Find(x, y int) (Point, bool) {
	for _, p := range m.points {
		if p.X == x && p.Y == y {
			return p, true
		}
	}
	return Point{}, false
}

func (m *Map) Collect(p Point) []Point {
	if !p.Broadcasting() {
		panic("cannot collect for non-broadcasting point")
	}
	out := []Point{}
	for _, po := range m.points {
		if po.Frequency == p.Frequency {
			out = append(out, po)
		}
	}
	return out
}

func ParseData(data []byte) (*Map, error) {
	out := Map{}
	r := bufio.NewReader(bytes.NewReader(data))

	y := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		line = strings.Trim(line, "\n")
		for x, char := range line {
			freq := string(char)
			if char == '.' {
				freq = ""
			} else {
				if !slices.Contains(out.uniqueFrequencies, freq) {
					out.uniqueFrequencies = append(out.uniqueFrequencies, freq)
				}
			}
			out.points = append(out.points, Point{
				X:         x,
				Y:         y,
				Frequency: freq,
			})
		}
		y++
	}

	return &out, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	m, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse data: %w", err)
	}

	antinodes := []Point{}
	for _, freq := range m.uniqueFrequencies {
		antinodes = append(antinodes, m.Antinodes(freq, 1)...)
	}
	antinodes = Unique(antinodes)

	fmt.Printf("Count: %d\n", len(antinodes))

	return nil
}

func PartTwo(ctx context.Context) error {
	m, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse data: %w", err)
	}

	antinodes := []Point{}
	for _, freq := range m.uniqueFrequencies {
		antinodes = append(antinodes, m.Antinodes(freq, int(math.Inf(1)))...)
	}
	antinodes = Unique(antinodes)

	fmt.Printf("Count: %d\n", len(antinodes))

	return nil
}
