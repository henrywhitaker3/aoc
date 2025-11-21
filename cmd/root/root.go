// Package root
package root

import (
	"github.com/henrywhitaker3/aoc/cmd/twentytwentyfive"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "aoc",
	}

	cmd.AddCommand(twentytwentyfive.Cmd())

	cmd.PersistentFlags().IntP("part", "p", 1, "The part of the day to run")

	return cmd
}
