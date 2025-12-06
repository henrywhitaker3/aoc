// Package day4
package day4

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
	X     int
	Y     int
	Paper bool
}

type Grid []Point

func (g Grid) MaxX() int {
	out := 0
	for _, p := range g {
		if p.X > out {
			out = p.X
		}
	}
	return out
}

func (g Grid) MaxY() int {
	out := 0
	for _, p := range g {
		if p.Y > out {
			out = p.Y
		}
	}
	return out
}

func (g Grid) Find(x, y int) (Point, bool) {
	for _, p := range g {
		if p.X == x && p.Y == y {
			return p, true
		}
	}
	return Point{}, false
}

func (g Grid) AdjacentPaper(p Point) int {
	out := 0
	for _, n := range g.adjacentPoints(p) {
		if n.Paper {
			out++
		}
	}
	return out
}

func (g Grid) adjacentPoints(p Point) []Point {
	movements := [][]int{
		{0, -1},  // Up
		{0, 1},   // Down
		{-1, 0},  // Left
		{1, 0},   // Right
		{-1, -1}, // Up+Left
		{-1, 1},  // Up+Right
		{1, -1},  // Down+Left
		{1, 1},   // Down+Right
	}
	out := []Point{}
	for _, m := range movements {
		if n, ok := g.Find(p.X+m[0], p.Y+m[1]); ok {
			out = append(out, n)
		}
	}
	return out
}

func (g Grid) MoveablePoints() []Point {
	out := []Point{}
	for _, p := range g {
		if !p.Paper {
			continue
		}
		if g.AdjacentPaper(p) < 4 {
			out = append(out, p)
		}
	}
	return out
}

func ParseData(data []byte) (Grid, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	out := Grid{}
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
		slog.Debug("found point", "x", x, "y", y, "char", string(char))
		out = append(out, Point{
			X:     x,
			Y:     y,
			Paper: char == '@',
		})
	}
	return out
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	grid, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Count: %d\n", len(grid.MoveablePoints()))

	return nil
}

func PartTwo(ctx context.Context) error {
	return nil
}
