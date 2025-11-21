package twentytwentyfive

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func dayOne() *cobra.Command {
	return &cobra.Command{
		Use: "01",
		RunE: func(cmd *cobra.Command, args []string) error {
			part, err := cmd.Flags().GetInt("part")
			if err != nil {
				return fmt.Errorf("get part: %w", err)
			}

			switch part {
			case 1:
				return partOne(cmd.Context())
			case 2:
				return partTwo(cmd.Context())
			default:
				return fmt.Errorf("invalid part")
			}
		},
	}
}

func partOne(ctx context.Context) error {
	return fmt.Errorf("not implemented yet")
}

func partTwo(ctx context.Context) error {
	return fmt.Errorf("not implemented yet")
}
