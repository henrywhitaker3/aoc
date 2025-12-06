// Package day6
package day6

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	partMatch = regexp.MustCompile(`([0-9-*.+]+)`)
)

type Operator string

const (
	Multiply = "*"
	Add      = "+"
)

type Operand[T any] struct {
	parsed T
}

func (o Operand[T]) Value() T {
	return o.parsed
}

type Calculation struct {
	Numbers  []Operand[int]
	Operator Operand[Operator]
}

func (c Calculation) Calculate() int {
	out := 0
	for i, n := range c.Numbers {
		if i == 0 {
			out = n.Value()
			continue
		}
		switch c.Operator.Value() {
		case Multiply:
			out *= n.Value()
		case Add:
			out += n.Value()
		}
	}
	return out
}

func SumResults(input []Calculation) int {
	out := 0
	for _, c := range input {
		out += c.Calculate()
	}
	return out
}

func ParseData(data []byte) ([]Calculation, error) {
	r := bufio.NewReader(bytes.NewReader(data))

	operands := [][]Operand[int]{}
	operators := []Operand[Operator]{}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		matches := partMatch.FindAllString(strings.Trim(line, "\n"), -1)
		switch matches[0] {
		case "+":
			fallthrough
		case "*":
			ops, err := parseOperators(matches)
			if err != nil {
				return nil, fmt.Errorf("parse operators: %w", err)
			}
			operators = append(operators, ops...)
		default:
			ops, err := parseDigits(matches)
			if err != nil {
				return nil, fmt.Errorf("parse digits: %w", err)
			}
			operands = append(operands, ops)
		}
	}

	out := []Calculation{}
	for i := range len(operands[0]) {
		calc := Calculation{
			Operator: operators[i],
		}
		for _, ops := range operands {
			calc.Numbers = append(calc.Numbers, ops[i])
		}
		out = append(out, calc)
	}

	return out, nil
}

func parseDigits(m []string) ([]Operand[int], error) {
	out := []Operand[int]{}
	for _, m := range m {
		dig, err := strconv.Atoi(m)
		if err != nil {
			return nil, fmt.Errorf("parse int: %w", err)
		}
		out = append(out, Operand[int]{
			parsed: dig,
		})
	}
	return out, nil
}

func parseOperators(m []string) ([]Operand[Operator], error) {
	out := []Operand[Operator]{}
	for _, m := range m {
		var op Operator
		switch m {
		case Multiply:
			op = Multiply
		case Add:
			op = Add
		default:
			return nil, fmt.Errorf("unknown operator %s", m)
		}
		out = append(out, Operand[Operator]{
			parsed: op,
		})
	}
	return out, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	calcs, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", SumResults(calcs))

	return nil
}

func PartTwo(ctx context.Context) error {
	return nil
}
