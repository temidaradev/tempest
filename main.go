package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	game "github.com/temidaradev/tempest/core"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowTitle("Tempest")
	ebiten.SetWindowSize(1280, 720)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
