package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/henrywhitaker3/aoc/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cmd := cmd.Cmd()
	cmd.SetContext(ctx)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
