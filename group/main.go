package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	auth := p.Group("/auth")
	auth.GET("/ping", func(ctx context.Context, pk *server.PianoKey) {
		pk.String(consts.StatusOK, "pong")
	})
	auth.GET("/binary", func(ctx context.Context, pk *server.PianoKey) {
		pk.String(consts.StatusOK, "lorain")
	})
	p.Play()
}
