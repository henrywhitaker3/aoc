// Package twentytwentyfive
package twentytwentyfive

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "2025",
	}

	cmd.AddCommand(dayOne())

	return cmd
}
