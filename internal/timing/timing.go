// Package timing
package timing

import (
	"log/slog"
	"time"
)

func Timed(f func()) {
	start := time.Now()
	f()
	slog.Info("ran function", "time", time.Since(start).String())
}
