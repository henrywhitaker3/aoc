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

	"github.com/henrywhitaker3/aoc/internal/caching"
	"github.com/henrywhitaker3/aoc/internal/grid"
	"github.com/henrywhitaker3/aoc/internal/timing"
)

type Point struct {
	x    int
	y    int
	char string
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

func (p Point) String() string {
	return p.char
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

func (p Point) Key() string {
	return fmt.Sprintf("%d,%d", p.X, p.y)
}

type Manifold struct {
	*grid.Grid[Point]
}

func NewManifold() *Manifold {
	return &Manifold{
		Grid: grid.New[Point](),
	}
}

func (m Manifold) Start() Point {
	for _, p := range m.Points() {
		if p.Start() {
			return p
		}
	}
	panic("not start point found")
}

func (m Manifold) Next(x, y int) (Point, bool) {
	_, p, ok := m.Find(x, y+1)
	return p, ok
}

func (m Manifold) Update(p Point, char string) {
	if !m.Grid.Update(p.X(), p.Y(), func(p Point) Point {
		p.char = char
		return p
	}) {
		panic("could not update point")
	}
}

func (m Manifold) String() string {
	out := []string{}
	y := 0
	str := ""
	for _, p := range m.Points() {
		if p.y > y {
			out = append(out, str)
			str = ""
			y++
		}
		str = fmt.Sprintf("%s%s", str, p.char)
	}
	return strings.Join(out, "\n")
}

func CountSplits(m *Manifold) int {
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
			next, ok := m.Next(beam.X(), beam.Y())
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
				if _, left, ok := m.Find(next.X()-1, next.Y()); ok {
					beams = append(beams, left)
				}
				if _, right, ok := m.Find(next.X()+1, next.Y()); ok {
					beams = append(beams, right)
				}
				break
			}
			beam.y = next.y
		}
		i++
		slog.Debug("beam escaped", "beam", beam)
	}

	fmt.Println(m.String())

	return len(splits)
}

func CountTimelines(m *Manifold) int {
	memo := caching.Memoise[int]()
	start := m.Start()
	return countTimelines(memo, m, start.X(), start.Y())
}

func countTimelines(memo *caching.Memoised[int], m *Manifold, x, y int) int {
	_, point, ok := m.Find(x, y)
	if !ok {
		return 1
	}

	call := func(m *Manifold, x, y int) func() int {
		return func() int {
			return countTimelines(memo, m, x, y)
		}
	}
	key := func(x, y int) string {
		return fmt.Sprintf("%d,%d", x, y)
	}

	if point.Split() {
		return memo.Run(
			key(point.X()-1, point.Y()+1),
			call(m, point.X()-1, point.Y()+1),
		) + memo.Run(
			key(point.X()+1, point.Y()+1),
			call(
				m,
				point.X()+1,
				point.Y()+1,
			),
		)
	}

	return memo.Run(key(point.X(), point.Y()+1), call(m, point.X(), point.Y()+1))
}

func indexOk[T any](s []T, i int) (T, bool) {
	if i < 0 || i >= len(s) {
		var out T
		return out, false
	}
	return s[i], true
}

func ParseData(data []byte) (*Manifold, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	out := NewManifold()
	y := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		out.Push(parseLine(strings.Trim(line, "\n"), y)...)
		y++
	}
	return out, nil
}

func parseLine(line string, y int) []Point {
	out := []Point{}
	for x, char := range line {
		out = append(out, Point{
			x:    x,
			y:    y,
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

	timing.Timed(func() {
		fmt.Printf("Splits: %d\n", CountSplits(man))
	})

	return nil
}

func PartTwo(ctx context.Context) error {
	man, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse data: %w", err)
	}
	slog.Debug("parsed data")

	timing.Timed(func() {
		fmt.Printf("Splits: %d\n", CountTimelines(man))
	})

	return nil
}
