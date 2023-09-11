package main

import (
	"context"
	"time"

	"github.com/B1NARY-GR0UP/inquisitor/core"
	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default(server.WithShutdownTimeout(3 * time.Second))
	p.OnRun = append(p.OnRun, func(ctx context.Context) error {
		core.Info("hello")
		return nil
	})
	p.OnRun = append(p.OnRun, func(ctx context.Context) error {
		core.Info("world")
		return nil
	})
	p.OnShutdown = append(p.OnShutdown, func(ctx context.Context) {
		core.Info("binary")
	})
	p.OnShutdown = append(p.OnShutdown, func(ctx context.Context) {
		core.Info("piano")
	})
	p.GET("/ping", func(ctx context.Context, pk *server.PianoKey) {
		pk.String(consts.StatusOK, "pong")
	})
	p.Play()
}
