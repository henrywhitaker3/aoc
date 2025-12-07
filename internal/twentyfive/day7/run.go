// Package day7
package day7

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type Point struct {
	X    int
	Y    int
	char string
}

func (p Point) Split() bool {
	return p.char == "^"
}

func (p Point) Start() bool {
	return p.char == "S"
}

func (p Point) Visited() bool {
	return p.char == "|"
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Manifold []Point

func (m Manifold) Start() Point {
	for _, p := range m {
		if p.Start() {
			return p
		}
	}
	panic("not start point found")
}

func (m Manifold) Find(x, y int) (int, Point, bool) {
	for i, p := range m {
		if p.X == x && p.Y == y {
			return i, p, true
		}
	}
	return 0, Point{}, false
}

func (m Manifold) Next(x, y int) (Point, bool) {
	_, p, ok := m.Find(x, y+1)
	return p, ok
}

func (m Manifold) Update(p Point, char string) {
	if i, _, ok := m.Find(p.X, p.Y); ok {
		if !p.Split() {
			m[i].char = char
		}
	}
}

func (m Manifold) String() string {
	out := []string{}
	y := 0
	str := ""
	for _, p := range m {
		if p.Y > y {
			out = append(out, str)
			str = ""
			y++
		}
		str = fmt.Sprintf("%s%s", str, p.char)
	}
	return strings.Join(out, "\n")
}

func CountSplits(m Manifold) int {
	splits := map[string]struct{}{}

	beams := []Point{m.Start()}
	i := 0
	for {
		beam, ok := indexOk(beams, i)
		if !ok {
			break
		}
		for {
			slog.Debug("checking beam", "beam", beam)
			next, ok := m.Next(beam.X, beam.Y)
			if !ok {
				// If it doesn't ecist, then we've exited the manifold
				break
			}
			if next.Visited() {
				break
			}
			m.Update(beam, "|")
			if next.Split() {
				// If the next one is a splitter, then check left and right
				// for new beams to add to the path
				slog.Debug("got split", "beam", beam, "split", next)
				splits[next.String()] = struct{}{}
				if _, left, ok := m.Find(next.X-1, next.Y); ok {
					beams = append(beams, left)
				}
				if _, right, ok := m.Find(next.X+1, next.Y); ok {
					beams = append(beams, right)
				}
				break
			}
			beam.Y = next.Y
		}
		i++
		slog.Debug("beam escaped", "beam", beam)
	}

	fmt.Println(m.String())

	return len(splits)
}

func indexOk[T any](s []T, i int) (T, bool) {
	if i < 0 || i >= len(s) {
		var out T
		return out, false
	}
	return s[i], true
}

func ParseData(data []byte) (Manifold, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	out := Manifold{}
	y := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		out = append(out, parseLine(strings.Trim(line, "\n"), y)...)
		y++
	}
	return out, nil
}

func parseLine(line string, y int) []Point {
	out := []Point{}
	for x, char := range line {
		out = append(out, Point{
			X:    x,
			Y:    y,
			char: string(char),
		})
	}
	return out
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	man, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse data: %w", err)
	}
	slog.Debug("parsed data")

	fmt.Printf("Splits: %d\n", CountSplits(man))

	return nil
}

func PartTwo(ctx context.Context) error {
	return nil
}
