package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	game "github.com/temidaradev/tempest/core"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
