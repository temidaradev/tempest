package main

import (
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/pkg/game"
)

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		slog.Error("error:", err)
	}
}