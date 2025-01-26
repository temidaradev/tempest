package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/ui"
)

type Game struct {
	tick    uint64
	paused  bool
	updated bool

	splash *ui.SplashView
}

func NewGame() *Game {
	g := &Game{
		tick:   0,
		splash: &ui.SplashView{},
	}

	return g
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.splash.Active() {
		g.splash.Draw(screen)
	}
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
