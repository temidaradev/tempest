package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/assets"
	"github.com/temidaradev/tempest/ui"
)

type Game struct {
	tick    uint64
	paused  bool
	updated bool

	splash *ui.SplashView
}

func NewGame() *Game {
	return &Game{
		tick:   0,
		splash: ui.NewSplashView(),
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.updated {
		if g.splash.Active() {
			g.splash.Draw(screen)
		}
	}

	screen.DrawImage(assets.Compact[0], &ebiten.DrawImageOptions{})
}

func (g *Game) Update() error {
	g.tick++

	if g.splash.Active() {
		g.splash.Update()
		g.updated = true
		return nil
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
