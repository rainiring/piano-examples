package main

import (
	"context"

	"github.com/B1NARY-GR0UP/piano/core/server"
	"github.com/B1NARY-GR0UP/piano/core/server/bin"
	"github.com/B1NARY-GR0UP/piano/pkg/consts"
)

func main() {
	p := bin.Default()
	p.GET("/query", func(ctx context.Context, pk *server.PianoKey) {
		pk.JSON(consts.StatusOK, server.M{
			"username": pk.Query("username"),
		})
	})
	p.Play()
}
