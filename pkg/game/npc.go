package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/assets"
)

type NPC struct {
	player *Player
	DIO    *ebiten.DrawImageOptions
	count  int
}

func (n *NPC) createNPC(screen *ebiten.Image) {
	frameCountIdle := 2
	i := (n.count / 45) % frameCountIdle
	n.DIO.GeoM.Scale(4, 4)
	n.DIO.GeoM.Translate(100, 335)
	cam.Draw(assets.ExBfTile[i], n.DIO, screen)
	n.DIO.GeoM.Reset()
}

func (n *NPC) Update() {
	n.count++
}
