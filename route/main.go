package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	// static route or common route
	p.GET("/ping", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"ping": "pong",
		})
	})
	// param route
	p.GET("/param/:username", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"username": pk.Param("username"),
		})
	})
	// wildcard route
	p.GET("/wild/*", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"route": "wildcard route",
		})
	})
	p.Play()
}
