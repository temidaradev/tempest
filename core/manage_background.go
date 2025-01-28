package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/assets"
)

func MakeBackground(screen *ebiten.Image) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			dio := &ebiten.DrawImageOptions{}
			dio.GeoM.Translate(2000-float64(i*50), 1900-float64(j*50))
			cam.Draw(assets.Coupe[1], dio, screen)
		}
	}
}
