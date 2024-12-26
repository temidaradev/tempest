package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/assets"
)

type NPC struct {
	DIO   *ebiten.DrawImageOptions
	count int
}

func (n *NPC) createNPC(screen *ebiten.Image) {
	frameCountIdle := 2
	i := (n.count / 45) % frameCountIdle
	n.DIO.GeoM.Scale(2, 2)
	n.DIO.GeoM.Translate(225, 330)
	cam.Draw(assets.ExBfTile[i], n.DIO, screen)
	n.DIO.GeoM.Reset()
}

func (n *NPC) Update() {
	n.count++
}
