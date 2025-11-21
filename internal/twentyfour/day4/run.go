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

type Graph []Point

func (g Graph) MaxX() int {
	x := 0
	for _, p := range g {
		if p.X > x {
			x = p.X
		}
	}
	return x
}

func (g Graph) MaxY() int {
	y := 0
	for _, p := range g {
		if p.Y > y {
			y = p.X
		}
	}
	return y
}

func (g Graph) Find(x, y int) (Point, bool) {
	for _, p := range g {
		if p.X == x && p.Y == y {
			return p, true
		}
	}
	return Point{}, false
}

type Point struct {
	X   int
	Y   int
	Val rune
}

func ParseData(data []byte) (Graph, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	out := Graph{}

	lines := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}

		for i, char := range strings.Trim(line, "\n") {
			out = append(out, Point{
				X:   i,
				Y:   lines,
				Val: char,
			})
		}
		lines++
	}

	return out, nil
}

var (
	directions = [][]int{
		{0, 1},   // Up
		{1, 0},   // Right
		{1, 1},   // Up+Right
		{1, -1},  // Down+Right
		{0, -1},  // Down
		{-1, 0},  // Left
		{-1, -1}, // Down+Left
		{-1, 1},  // Up+Left
	}
)

func CountXmas(g Graph) (int, error) {
	starters := []Point{}
	for _, p := range g {
		if p.Val == 'X' {
			starters = append(starters, p)
		}
	}

	count := 0
	for _, p := range starters {
		for _, d := range directions {
			if getWord(g, p, d) {
				count++
			}
		}
	}

	return count, nil
}

func getWord(g Graph, p Point, dir []int) bool {
	search := []rune{'M', 'A', 'S'}
	for i := range 3 {
		next, ok := g.Find(p.X+dir[0], p.Y+dir[1])
		if !ok {
			return false
		}
		p = next
		if next.Val != search[i] {
			slog.Debug("got match")
			return false
		}
	}
	return true
}

var (
	crossDirections = [][][]int{
		{{1, 1}, {-1, -1}}, // /
		{{-1, 1}, {1, -1}}, // \
	}
)

func CountCrossMas(g Graph) (int, error) {
	starters := []Point{}
	for _, p := range g {
		if p.Val == 'A' {
			starters = append(starters, p)
		}
	}

	count := 0
	for _, p := range starters {
		diags, ok := getDiagonals(g, p)
		if !ok {
			continue
		}
		if (diags[0] == "MAS" || diags[0] == "SAM") && (diags[1] == "MAS" || diags[1] == "SAM") {
			count++
		}
	}

	return count, nil
}

func getDiagonals(g Graph, p Point) ([]string, bool) {
	out := []string{}
	for _, d := range crossDirections {
		first, ok := g.Find(p.X+d[0][0], p.Y+d[0][1])
		if !ok {
			return nil, false
		}
		second, ok := g.Find(p.X+d[1][0], p.Y+d[1][1])
		if !ok {
			return nil, false
		}
		out = append(out, fmt.Sprintf("%vA%v", string(first.Val), string(second.Val)))
	}
	return out, true
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	graph, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	count, err := CountXmas(graph)
	if err != nil {
		return fmt.Errorf("count xmas: %w", err)
	}

	fmt.Printf("Count: %d\n", count)

	return nil
}

func PartTwo(ctx context.Context) error {
	graph, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	count, err := CountCrossMas(graph)
	if err != nil {
		return fmt.Errorf("count crossmas: %w", err)
	}

	fmt.Printf("Count: %d\n", count)

	return nil
}
