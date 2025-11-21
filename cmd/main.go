// Package cmd
package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/henrywhitaker3/aoc/internal/twentyfive"
	"github.com/henrywhitaker3/aoc/internal/twentyfour"
	"github.com/spf13/cobra"
)

func init() {
	twentyfour.Register(solutions)
	twentyfive.Register(solutions)
}

var (
	log bool
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "aoc",
		PreRun: func(cmd *cobra.Command, args []string) {
			var out io.Writer = &bytes.Buffer{}
			if log {
				out = os.Stdout
			}
			slog.SetDefault(slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			})))
		},
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

	cmd.Flags().BoolVarP(&log, "log", "l", false, "Whether to show log output")

	return cmd
}

var (
	solutions = Solutions{}
)

type Solutions map[string]func(context.Context) error

func (s Solutions) key(year int, day int, part int) string {
	return fmt.Sprintf("%d:%d:%d", year, day, part)
}

func (s Solutions) Get(year int, day int, part int) (func(context.Context) error, bool) {
	f, ok := s[s.key(year, day, part)]
	return f, ok
}

func (s Solutions) Set(year, day, part int, f func(context.Context) error) {
	s[s.key(year, day, part)] = f
}
