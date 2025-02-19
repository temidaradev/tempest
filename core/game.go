package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/setanarut/kamera/v2"
	"github.com/temidaradev/esset/v2"
	"github.com/temidaradev/tempest/assets"
	"github.com/temidaradev/tempest/ui"
)

type Game struct {
	tick    uint64
	updated bool

	car    *Car
	menu   *Menu
	splash *ui.SplashView
}

func init() {
	assets.FontFace, _ = esset.GetFont(assets.MyFont, 48)

	cam.SmoothType = kamera.SmoothDamp
	cam.SmoothOptions.SmoothDampTimeY = 0.5
}

func NewGame() *Game {
	cam.ShakeEnabled = true
	g := &Game{
		menu:   &Menu{},
		splash: ui.NewSplashView(),
	}
	g.car = &Car{
		X:     0,
		Y:     0,
		DIO:   &ebiten.DrawImageOptions{},
		angle: 0,
		speed: 0,
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
		if g.menu.sDone {
			screen.Clear()
			MakeBackground(screen)
			g.car.DrawCar(screen, g.menu.IndexC, int(g.menu.index))
		}
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
		if g.menu.sDone {
			cam.LookAt(g.car.X, g.car.Y)
			g.car.Update()
			if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				return ebiten.Termination
			}
		}
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
