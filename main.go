package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/pkg/game"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(1920, 1080)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
