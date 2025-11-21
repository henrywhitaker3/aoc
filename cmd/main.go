// Package cmd
package cmd

import (
	"context"
	"fmt"

	twentyfivedayone "github.com/henrywhitaker3/aoc/internal/twentyfive/dayone"
	twentyfourdayone "github.com/henrywhitaker3/aoc/internal/twentyfour/dayone"
	"github.com/spf13/cobra"
)

func init() {
	solutions.Set(2025, 1, 1, twentyfivedayone.PartOne)
	solutions.Set(2025, 1, 2, twentyfivedayone.PartTwo)
	solutions.Set(2024, 1, 1, twentyfourdayone.PartOne)
	solutions.Set(2024, 1, 2, twentyfourdayone.PartTwo)
}

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "aoc",
		RunE: func(cmd *cobra.Command, args []string) error {
			year, err := cmd.Flags().GetInt("year")
			if err != nil {
				return fmt.Errorf("get year: %w", err)
			}
			day, err := cmd.Flags().GetInt("day")
			if err != nil {
				return fmt.Errorf("get day: %w", err)
			}
			part, err := cmd.Flags().GetInt("part")
			if err != nil {
				return fmt.Errorf("get part: %w", err)
			}

			f, ok := solutions.Get(year, day, part)
			if !ok {
				return fmt.Errorf("no entry for specified solution")
			}

			return f(cmd.Context())
		},
	}

	cmd.Flags().IntP("year", "y", 2025, "The year of the day to run")
	cmd.Flags().IntP("day", "d", 0, "The day to run")
	cmd.Flags().IntP("part", "p", 0, "The part of the day to run")

	return cmd
}

var (
	solutions = Solutions{}
)

type PartFunc func(context.Context) error

type Solutions map[string]PartFunc

func (s Solutions) key(year int, day int, part int) string {
	return fmt.Sprintf("%d:%d:%d", year, day, part)
}

func (s Solutions) Get(year int, day int, part int) (PartFunc, bool) {
	f, ok := s[s.key(year, day, part)]
	return f, ok
}

func (s Solutions) Set(year, day, part int, f PartFunc) {
	s[s.key(year, day, part)] = f
}
