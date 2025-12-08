// Package grid
package grid

import (
	"fmt"

	"github.com/henrywhitaker3/aoc/internal/caching"
)

type GridMember interface {
	X() int
	Y() int
	String() string
}

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

func (p Point) String() string {
	return "."
}

type Grid[T GridMember] struct {
	points []T
	search *caching.Memoised[gridSearch[T]]
}

func New[T GridMember]() *Grid[T] {
	return &Grid[T]{
		points: []T{},
		search: caching.Memoise[gridSearch[T]](),
	}
}

func (g *Grid[T]) Push(p ...T) {
	g.points = append(g.points, p...)
}

type gridSearch[T GridMember] struct {
	index int
	point T
	ok    bool
}

func (g *Grid[T]) Find(x, y int) (int, T, bool) {
	key := fmt.Sprintf("%d,%d", x, y)
	res := g.search.Run(key, func() gridSearch[T] {
		return g.find(x, y)
	})
	return res.index, res.point, res.ok
}

func (g *Grid[T]) find(x, y int) gridSearch[T] {
	for i, p := range g.points {
		if p.X() == x && p.Y() == y {
			return gridSearch[T]{
				index: i,
				point: p,
				ok:    true,
			}
		}
	}
	return gridSearch[T]{}
}

func (g *Grid[T]) Update(x, y int, f func(T) T) bool {
	i, point, ok := g.Find(x, y)
	if !ok {
		return false
	}
	g.points[i] = f(point)
	g.search.Replace(fmt.Sprintf("%d,%d", x, y), gridSearch[T]{
		index: i,
		point: g.points[i],
		ok:    true,
	})
	return true
}

func (g Grid[T]) Points() []T {
	return g.points
}
