package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/henrywhitaker3/aoc/cmd/root"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cmd := root.Cmd()
	cmd.SetContext(ctx)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
