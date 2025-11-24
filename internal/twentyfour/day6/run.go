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
	"runtime"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Point struct {
	X        int
	Y        int
	Blocking bool
	Visited  int
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
		if p.Visited > 0 {
			m.guard.X = p.X
			m.guard.Y = p.Y
		}
	}

	return m
}

func (m *Map) GuardPOS() (int, Point, bool) {
	return m.Find(m.guard.X, m.guard.Y)
}

func (m *Map) Clone() *Map {
	cloned := &Map{
		points: slices.Clone(m.points),
		maxX:   m.maxX,
		maxY:   m.maxY,
		guard: Point{
			X: m.guard.X,
			Y: m.guard.Y,
		},
		direction: m.direction,
	}
	return cloned
}

func (m *Map) Find(x, y int) (int, Point, bool) {
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

// Move returns true if the guard moved to a new position, false if off the map.
// Also returns a second bool that shows if the guard has already been to that
// position before.
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
	m.points[i].Visited++

	return true
}

func (m *Map) SumMoves() int {
	ok := m.Move()
	for ok {
		slog.Debug("visited position")
		// Carry on moving until we don't anymore
		ok = m.Move()
	}

	out := 0
	for _, p := range m.points {
		if p.Visited > 0 {
			out++
		}
	}

	return out
}

func SumLoops(ctx context.Context, points []Point) int {
	blockable := []Point{}
	for _, p := range points {
		if p.Visited == 0 && !p.Blocking {
			blockable = append(blockable, p)
			slog.Debug("got block point", "x", p.X, "y", p.Y)
		}
	}

	loops := &atomic.Int64{}
	ch := make(chan Point, runtime.NumCPU()*2)
	wg := &sync.WaitGroup{}
	for range runtime.NumCPU() {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case b := <-ch:
					if determineIfLoop(b, points) {
						loops.Add(1)
					}
					wg.Done()
				}
			}
		}()
	}

	for _, b := range blockable {
		ch <- b
		wg.Add(1)
	}

	wg.Wait()

	return int(loops.Load())
}

func determineIfLoop(b Point, points []Point) bool {
	slog.Debug("trying block point", "x", b.X, "y", b.Y)
	cloned := NewMap(slices.Clone(points))
	i, _, _ := cloned.Find(b.X, b.Y)
	cloned.points[i].Blocking = true

	moves := 0
	dupes := 0
	ok := cloned.Move()
	for ok {
		moves++
		_, pos, _ := cloned.GuardPOS()
		if pos.Visited > 1 {
			slog.Debug("we've been here before")
			dupes++
		}
		// Guessing that if we have been to the same 20 spots, we're in a loop
		if dupes > 1000 {
			slog.Debug("we looped")
			return true
		}
		ok = cloned.Move()
	}
	slog.Debug("we moved", "moves", moves)
	return false
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
				p.Visited = 1
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
	start := time.Now()

	m, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum positions: %d\n", m.SumMoves())

	slog.Debug("found moves", "taken", time.Since(start).String())

	return nil
}

func PartTwo(ctx context.Context) error {
	start := time.Now()

	m, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum loops: %d\n", SumLoops(ctx, m.points))

	slog.Debug("found loops", "taken", time.Since(start).String())

	return nil
}
