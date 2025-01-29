package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	game "github.com/temidaradev/tempest/core"
)

func main() {
	g := game.NewGame()
	fps := ebiten.ActualFPS()
	title := fmt.Sprint("Tempest FPS: %v", fps)
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
