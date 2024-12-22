package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/ehh24/assets"
)

type Player struct {
	X, Y  float64
	DIO   *ebiten.DrawImageOptions
	speed float64
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.DIO.GeoM.Scale(1, 1)
	p.DIO.GeoM.Translate(p.X-playerOffsetX+50, p.Y-playerOffsetY+50)
	cam.Draw(assets.IdleTile[0], p.DIO, screen)
	p.DIO.GeoM.Reset()
}

func (p *Player) Update() {
	x, y := Axis()
	p.X += x * p.speed
	p.Y += y * p.speed
}
