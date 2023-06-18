package main

import (
	"context"
	"os/signal"
	"syscall"

	"fishing-store/internal/app"
	"fishing-store/internal/app/config"
)

func main() {
	cfg := config.NewConfig()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app.Run(ctx, cfg)
}
