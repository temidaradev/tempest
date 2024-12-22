package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/setanarut/kamera/v2"
	"github.com/temidaradev/ehh24/assets"
)

var (
	w, h             = 635., 475.
	groundY          = 395.
	unit             = 10.
	targetX, targetY = w / 2, h / 2
	cam              = kamera.NewCamera(targetX, targetY, w, h)
)

type Game struct {
	player *Player
}

func NewGame() *Game {
	g := &Game{}
	cam.ShakeEnabled = true

	g.player = &Player{
		X:     0,
		Y:     0,
		DIO:   &ebiten.DrawImageOptions{},
		speed: 10,
	}

	return g
}

var playerOffsetX = float64(assets.IdleTile[0].Bounds().Dx() / 2)
var playerOffsetY = float64(assets.IdleTile[0].Bounds().Dy() / 2)

func (g *Game) Draw(screen *ebiten.Image) {
	DIO := &ebiten.DrawImageOptions{}
	DIO.GeoM.Scale(0.5, 0.5)
	cam.Draw(assets.CityNight, DIO, screen)
	g.player.Draw(screen)
}

func (g *Game) Update() error {
	g.player.Update()
	cam.LookAt(g.player.X, g.player.Y)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func Axis() (axisX, axisY float64) {
	// if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	axisY -= 1
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyS) {
	// 	axisY += 1
	// }
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if axisY == groundY*unit {
			axisY = -10 * unit
		}

		if axisY > groundY*unit {
			axisY = groundY * unit
		}
		if axisY < 20*unit {
			axisY -= 8
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		axisX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		axisX += 1
	}
	return axisX, axisY
}
