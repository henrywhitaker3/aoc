package common

import (
	"log/slog"
	"testing"
)

type testWriter struct {
	t testing.TB
}

func (t testWriter) Write(data []byte) (int, error) {
	t.t.Log(string(data))
	return len(data), nil
}

func TestLogger(t testing.TB) {
	slog.SetDefault(slog.New(slog.NewJSONHandler(&testWriter{t: t}, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))
}
