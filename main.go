package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/pkg/game"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
