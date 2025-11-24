// Package day7
package day7

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type Equation struct {
	Answer int
	Inputs []int
}

func Evaluate(ctx context.Context, equations []Equation, concat bool) int {
	out := &atomic.Int64{}

	ch := make(chan Equation, runtime.NumCPU()*2)
	wg := &sync.WaitGroup{}

	for range runtime.NumCPU() {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case job := <-ch:
					processEquation(job, wg, out, concat)
				}
			}
		}()
	}

	for _, eq := range equations {
		ch <- eq
		wg.Add(1)
	}

	wg.Wait()

	return int(out.Load())
}

func processEquation(eq Equation, wg *sync.WaitGroup, out *atomic.Int64, concatenate bool) {
	defer wg.Done()
	if evaluate(eq.Answer, eq.Inputs, 0, concatenate) {
		out.Add(int64(eq.Answer))
	}
}

func evaluate(expected int, inputs []int, sum int, concat bool) bool {
	if len(inputs) == 0 {
		return sum == expected
	}

	if sum > expected {
		return false
	}

	if evaluate(expected, inputs[1:], sum+inputs[0], concat) {
		return true
	}

	if concat && evaluate(expected, inputs[1:], concatenate(sum, inputs[0]), concat) {
		return true
	}

	return evaluate(expected, inputs[1:], sum*inputs[0], concat)
}

func concatenate(first, second int) int {
	out, err := strconv.Atoi(fmt.Sprintf("%d%d", first, second))
	if err != nil {
		panic(err)
	}
	return out
}

func ParseData(data []byte) ([]Equation, error) {
	out := []Equation{}

	r := bufio.NewReader(bytes.NewReader(data))

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read line: %w", err)
		}
		line = strings.Trim(line, "\n")
		eq, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parse line: %w", err)
		}
		out = append(out, eq)
	}

	return out, nil
}

func parseLine(line string) (Equation, error) {
	out := Equation{}

	spl := strings.Split(line, ": ")

	answer, err := strconv.Atoi(spl[0])
	if err != nil {
		return out, fmt.Errorf("parse answer: %w", err)
	}

	out.Answer = answer

	spl = strings.Split(spl[1], " ")

	for _, s := range spl {
		num, err := strconv.Atoi(s)
		if err != nil {
			return out, fmt.Errorf("parse input %s: %w", s, err)
		}
		out.Inputs = append(out.Inputs, num)
	}

	return out, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	eqs, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", Evaluate(ctx, eqs, false))

	return nil
}

func PartTwo(ctx context.Context) error {
	eqs, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	fmt.Printf("Sum: %d\n", Evaluate(ctx, eqs, true))

	return nil
}
