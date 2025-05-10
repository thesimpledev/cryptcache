package main

import (
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	}))

	if err := parse(); err != nil {
		logger.Error("Command failed", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}
