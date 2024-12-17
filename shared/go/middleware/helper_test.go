package middleware

import (
	"bytes"
	"log/slog"
)

func makeLogger() *slog.Logger {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	return logger
}

func makeLoggerAndBuffer() (*slog.Logger, *bytes.Buffer) {
	buf := bytes.Buffer{}
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	return logger, &buf
}
