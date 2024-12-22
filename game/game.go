package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/setanarut/kamera/v2"
	"github.com/temidaradev/ehh24/assets"
)

var (
	w, h             = 635., 475.
	groundY          = 400.
	unit             = 10.
	targetX, targetY = w / 2, h / 2
	cam              = kamera.NewCamera(targetX, targetY, w, h)
)

type Game struct {
	player *Player
}

func init() {
	cam.SmoothType = kamera.SmoothDamp
	cam.SmoothOptions.SmoothDampTimeY = 0.2
}

func NewGame() *Game {
	g := &Game{}
	cam.ShakeEnabled = true

	g.player = &Player{
		DIO:   &ebiten.DrawImageOptions{},
		speed: 10,
	}

	return g
}

var playerOffsetX = float64(assets.IdleTile[0].Bounds().Dx() / 2)
var playerOffsetY = float64(assets.IdleTile[0].Bounds().Dy() / 2)

func (g *Game) Draw(screen *ebiten.Image) {
	HandleBackground(screen)
	g.player.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %v Y:%v", g.player.X, g.player.Y))
}

func (g *Game) Update() error {
	g.player.Update()
	if g.player.X >= 0 && g.player.X <= 1660 {
		cam.LookAt(g.player.X, g.player.Y)
	} else if g.player.X >= 1660 {
		cam.LookAt(1660, g.player.Y)
	} else if g.player.X <= 0 {
		cam.LookAt(0, g.player.Y)
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
