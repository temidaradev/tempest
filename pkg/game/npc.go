package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/esset/v2"
	"github.com/temidaradev/tempest/assets"
)

type NPC struct {
	X, Y  float64
	DIO   *ebiten.DrawImageOptions
	count int
}

func (n *NPC) createNPC(screen *ebiten.Image) {
	n.X = 100
	n.Y = 335
	frameCountIdle := 2
	i := (n.count / 45) % frameCountIdle
	n.DIO.GeoM.Scale(4, 4)
	n.DIO.GeoM.Translate(n.X, n.Y)
	cam.Draw(assets.ExBfTile[i], n.DIO, screen)
	esset.DrawText(screen, assets.MyFont, "Mecit", 16, n.X, n.Y, color.White)
	n.DIO.GeoM.Reset()
}

func (n *NPC) Update() {
	n.count++
}
