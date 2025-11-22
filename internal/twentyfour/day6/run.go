// Package day6
package day6

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
	X        int
	Y        int
	Blocking bool
	Visited  bool
}

type Map struct {
	points []Point

	maxX int
	maxY int

	guard Point

	direction []int
}

func NewMap(points []Point) *Map {
	m := &Map{
		points:    points,
		direction: []int{0, -1},
		guard:     Point{},
	}

	for _, p := range m.points {
		if p.X > m.maxX {
			m.maxX = p.X
		}
		if p.Y > m.maxY {
			m.maxY = p.Y
		}
		if p.Visited {
			m.guard.X = p.X
			m.guard.Y = p.Y
		}
	}

	return m
}

func (m *Map) Find(x, y int) (int, Point, bool) {
	slog.Debug("finding pos", "x", x, "y", y)
	if x < 0 || x > m.maxX || y < 0 || y > m.maxY {
		return 0, Point{}, false
	}

	for i, p := range m.points {
		if p.X == x && p.Y == y {
			return i, p, true
		}
	}

	panic(fmt.Errorf("could not find position when it should be there %d,%d", x, y))
}

func (m *Map) next() (int, Point, bool) {
	return m.Find(m.guard.X+m.direction[0], m.guard.Y+m.direction[1])
}

// Move returns true if the guard moved to a new position, false if off the map
func (m *Map) Move() bool {
	i, next, ok := m.next()
	if !ok {
		return false
	}

	if next.Blocking {
		newDir := changeDirection(m.direction)
		m.direction[0] = newDir[0]
		m.direction[1] = newDir[1]
		return m.Move()
	}

	m.guard.X = next.X
	m.guard.Y = next.Y
	m.points[i].Visited = true

	return true
}

func (m *Map) SumMoves() int {
	for m.Move() {
		slog.Debug("visited position")
		// Carry on moving until we don't anymore
	}

	out := 0
	for _, p := range m.points {
		if p.Visited {
			out++
		}
	}

	return out
}

func ParseData(data []byte) (*Map, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	points := []Point{}
	lines := 0

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
			p := Point{
				X: x,
				Y: lines,
			}
			switch char {
			case '#':
				p.Blocking = true
			case '^':
				p.Visited = true
			}
			points = append(points, p)
		}
		lines++
	}

	return NewMap(points), nil
}

func changeDirection(current []int) []int {
	if current[0] == 0 && current[1] == -1 {
		slog.Debug("moving right now")
		return []int{1, 0}
	}
	if current[0] == 1 && current[1] == 0 {
		slog.Debug("moving down now")
		return []int{0, 1}
	}
	if current[0] == 0 && current[1] == 1 {
		slog.Debug("moving left now")
		return []int{-1, 0}
	}
	if current[0] == -1 && current[1] == 0 {
		slog.Debug("moving up now")
		return []int{0, -1}
	}
	panic(fmt.Errorf("unhandled direction %d,%d", current[0], current[1]))
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	m, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum positions: %d\n", m.SumMoves())

	return nil
}

func PartTwo(ctx context.Context) error {
	return nil
}
