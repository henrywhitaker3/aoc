// Package day1
package day1

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
)

type Turn struct {
	Direction Direction
	Quantity  int
}

func ParseData(data []byte) ([]Turn, error) {
	out := []Turn{}
	r := bufio.NewReader(bytes.NewReader(data))

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		turn, err := parseLine(strings.Trim(line, "\n"))
		if err != nil {
			return nil, fmt.Errorf("parse line: %w", err)
		}
		out = append(out, turn)
	}

	return out, nil
}

func parseLine(line string) (Turn, error) {
	dir := string(line[0])
	quantRaw := line[1:]
	quant, err := strconv.Atoi(quantRaw)
	if err != nil {
		return Turn{}, fmt.Errorf("parse quantity %s: %w", quantRaw, err)
	}
	return Turn{
		Direction: Direction(dir),
		Quantity:  quant,
	}, nil
}

type Dial struct {
	pos int
}

func NewDial(start int) *Dial {
	return &Dial{
		pos: start,
	}
}

func (d *Dial) Go(turns []Turn, final func(pos int), click func(pos int)) {
	for _, t := range turns {
		d.move(t.Direction, t.Quantity, click)
		final(d.pos)
	}
}

func (d *Dial) move(direction Direction, amount int, cb func(pos int)) {
	for range amount {
		switch direction {
		case Left:
			upd := d.pos - 1
			if upd < 0 {
				upd = 99
			}
			d.pos = upd
		case Right:
			upd := d.pos + 1
			if upd == 100 {
				upd = 0
			}
			d.pos = upd
		default:
			panic("invalid direction")
		}
		cb(d.pos)
	}
}

func CountZeroes(turns []Turn, start int) int {
	dial := NewDial(start)
	count := 0
	dial.Go(turns, func(pos int) {
		if pos == 0 {
			count++
		}
	}, func(pos int) {})
	return count
}

func CountZeroClicks(turns []Turn, start int) int {
	dial := NewDial(start)
	count := 0
	dial.Go(turns, func(pos int) {}, func(pos int) {
		if pos == 0 {
			count++
		}
	})
	return count
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	turns, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Password: %d\n", CountZeroes(turns, 50))

	return nil
}

func PartTwo(ctx context.Context) error {
	turns, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Password: %d\n", CountZeroClicks(turns, 50))

	return nil
}
