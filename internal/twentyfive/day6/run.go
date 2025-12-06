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
	"slices"
	"strconv"
	"strings"
)

var (
	partMatch = regexp.MustCompile(`([0-9]+)|([+*]{1})`)
)

type Operator string

const (
	Multiply Operator = "*"
	Add      Operator = "+"
)

type Operand int

func (o Operand) Digits() int {
	out := 0
	num := o
	for num > 0 {
		num = num / 10
		out++
	}
	return out
}

type Calculation struct {
	Numbers  []Operand
	Operator Operator
}

func (c Calculation) Calculate() int {
	out := 0
	for i, n := range c.Numbers {
		if i == 0 {
			out = int(n)
			continue
		}
		switch c.Operator {
		case Multiply:
			out *= int(n)
		case Add:
			out += int(n)
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

	operands := [][]Operand{}
	operators := []Operator{}
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

func ParseRTL(data []byte) ([]Calculation, error) {
	rotated := rotate(data)
	spl := strings.Split(rotated, "\n")

	unparsed := [][]string{}
	buffer := []string{}
	for _, s := range spl {
		matches := partMatch.FindAllString(s, -1)
		if len(matches) == 0 {
			// We've reached the end of the calc
			unparsed = append(unparsed, slices.Clone(buffer))
			buffer = []string{}
			continue
		}
		buffer = append(buffer, matches...)
	}
	unparsed = append(unparsed, buffer)

	out := []Calculation{}
	for _, group := range unparsed {
		calc := Calculation{}
		for _, m := range group {
			switch m[0] {
			case '+':
				calc.Operator = Add
				continue
			case '*':
				calc.Operator = Multiply
				continue
			}
			dig, err := strconv.Atoi(m)
			if err != nil {
				return nil, fmt.Errorf("parse digit: %w", err)
			}
			calc.Numbers = append(calc.Numbers, Operand(dig))
		}
		slices.Reverse(calc.Numbers)
		out = append(out, calc)
	}

	slices.Reverse(out)

	return out, nil
}

func rotate(data []byte) string {
	lines := strings.Split(string(data), "\n")

	rows := len(lines)
	cols := 0
	for _, line := range lines {
		cols = max(cols, len(line))
	}

	rotated := make([]string, cols)

	for i := range rows {
		line := lines[i]
		for j := 0; j < cols; j++ {
			if j < len(line) {
				rotated[j] += string(line[j])
			}
		}
	}

	return strings.Join(rotated, "\n")
}

func parseDigits(m []string) ([]Operand, error) {
	out := []Operand{}
	for _, m := range m {
		if m == "" {
			continue
		}
		dig, err := strconv.Atoi(m)
		if err != nil {
			return nil, fmt.Errorf("parse int: %w", err)
		}
		out = append(out, Operand(dig))
	}
	return out, nil
}

func parseOperators(m []string) ([]Operator, error) {
	out := []Operator{}
	for _, m := range m {
		if m == "" {
			continue
		}
		var op Operator
		switch Operator(m) {
		case Multiply:
			op = Multiply
		case Add:
			op = Add
		default:
			return nil, fmt.Errorf("unknown operator %s", m)
		}
		out = append(out, op)
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
	calcs, err := ParseRTL([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", SumResults(calcs))

	return nil
}
