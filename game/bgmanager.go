package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/ehh24/assets"
)


func HandleBackground(screen *ebiten.Image) {
	DIO := &ebiten.DrawImageOptions{}
	DIO.GeoM.Scale(1, 1)
	DIO.GeoM.Translate(0, -815)
	cam.Draw(assets.CityNight, DIO, screen)
}
