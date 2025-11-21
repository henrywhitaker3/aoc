// Package day5
package day5

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Rules []Rule

func (r Rules) GetCorrectUpdates(updates []Update) []Update {
	out := []Update{}
ol:
	for _, u := range updates {
		for _, r := range r {
			if !r.Passes(u) {
				continue ol
			}
		}
		out = append(out, u)
	}

	return out
}

func (r Rules) GetCorrectedUpdates(updates []Update) []Update {
	out := []Update{}
ol:
	for _, u := range updates {
		for _, rule := range r {
			if !rule.Passes(u) {
				out = append(out, u)
				continue ol
			}
		}
	}

	for len(out) != len(r.GetCorrectUpdates(out)) {
		out = r.correctUpdates(out)
	}

	return out
}

func (r Rules) correctUpdates(u []Update) []Update {
	out := []Update{}

	for _, u := range u {
		u = slices.Clone(u)
		for _, r := range r {
			for !r.Passes(u) {
				slog.Debug("correcting value", "rule", r, "update", u)
				u = correctUpdate(r, u)
			}
		}
		slog.Debug("got corrected value", "update", u)
		out = append(out, u)
	}

	return out
}

func correctUpdate(rule Rule, u Update) Update {
	leftPos := slices.Index(u, rule.Left)
	leftVal := u[leftPos]
	rightPos := slices.Index(u, rule.Right)
	rightVal := u[rightPos]

	u[leftPos] = rightVal
	u[rightPos] = leftVal

	return u
}

type Rule struct {
	Left  int
	Right int
}

func (r Rule) Passes(updates Update) bool {
	// Pass rules that don't contain both
	if !slices.Contains(updates, r.Left) || !slices.Contains(updates, r.Right) {
		return true
	}

	leftPos := slices.Index(updates, r.Left)
	rightPos := slices.Index(updates, r.Right)

	return leftPos < rightPos
}

type Update []int

func (u Update) getMiddleValue() int {
	return u[int64(math.Floor(float64(len(u))/2))]
}

func sumMiddleValues(updates []Update) int {
	out := 0
	for _, u := range updates {
		out += u.getMiddleValue()
	}
	return out
}

func ParseData(data []byte) (Rules, []Update, error) {
	spl := strings.Split(string(data), "\n\n")
	rulesRaw := strings.Split(spl[0], "\n")
	updatesRaw := strings.Split(spl[1], "\n")

	rules, err := parseRules(rulesRaw)
	if err != nil {
		return nil, nil, fmt.Errorf("parse rules: %w", err)
	}

	updates, err := parseUpdates(updatesRaw)
	if err != nil {
		return nil, nil, fmt.Errorf("parse updates: %w", err)
	}

	return rules, updates, nil
}

func parseRules(input []string) (Rules, error) {
	out := Rules{}

	for _, rule := range input {
		if rule == "" {
			continue
		}
		spl := strings.Split(rule, "|")

		left, err := strconv.Atoi(spl[0])
		if err != nil {
			return nil, fmt.Errorf("parse left page as int: %w", err)
		}
		right, err := strconv.Atoi(spl[1])
		if err != nil {
			return nil, fmt.Errorf("parse right page as int: %w", err)
		}
		out = append(out, Rule{
			Left:  left,
			Right: right,
		})
	}

	return out, nil
}

func parseUpdates(input []string) ([]Update, error) {
	out := []Update{}

	for _, update := range input {
		if update == "" {
			continue
		}
		spl := strings.Split(update, ",")
		item := Update{}
		for _, s := range spl {
			val, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("parse int: %w", err)
			}
			item = append(item, val)
		}
		out = append(out, item)
	}

	return out, nil
}

var (
	//go:embed input.txt
	input string
)

func PartOne(ctx context.Context) error {
	rules, updates, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	correct := rules.GetCorrectUpdates(updates)

	fmt.Printf("Sum: %d\n", sumMiddleValues(correct))

	return nil
}

func PartTwo(ctx context.Context) error {
	rules, updates, err := ParseData([]byte(input))
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}

	correct := rules.GetCorrectedUpdates(updates)

	fmt.Printf("Sum: %d\n", sumMiddleValues(correct))

	return nil
}
