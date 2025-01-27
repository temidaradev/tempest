package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/ui"
)

type Game struct {
	tick    uint64
	paused  bool
	updated bool

	menu   *Menu
	splash *ui.SplashView
}

func NewGame() *Game {
	g := &Game{
		menu:   &Menu{},
		splash: ui.NewSplashView(),
	}
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.updated {
		if g.splash.Active() {
			g.splash.Draw(screen)
		}
	} else {
		g.menu.SelectCar(screen)
	}
}

func (g *Game) Update() error {
	g.tick++

	if g.splash.Active() {
		g.splash.Update()
		g.updated = true
		return nil
	} else {
		g.updated = false
		g.menu.HandleSelectCar()
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
