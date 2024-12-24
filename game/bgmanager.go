package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/assets"
)

func HandleBackground(screen *ebiten.Image) {
	DIO := &ebiten.DrawImageOptions{}
	DIO.GeoM.Scale(1, 1)
	DIO.GeoM.Translate(0, -815)
	cam.Draw(assets.CityNight, DIO, screen)
}

func HandleInterior(screen *ebiten.Image) {
	DIO := &ebiten.DrawImageOptions{}
	DIO.GeoM.Scale(0.3, 0.32)
	DIO.GeoM.Translate(-50, 0)
	cam.Draw(assets.Interior, DIO, screen)
}
