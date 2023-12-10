package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/nakaaaa/lamend"
)

func main() {
	ctx := context.Background()

	exit, err := lamend.Start(ctx)
	if err != nil {
		slog.Warn(err.Error())
	}
	os.Exit(exit)
}
