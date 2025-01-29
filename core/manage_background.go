package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/assets"
)

func MakeBackground(screen *ebiten.Image) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			dio := &ebiten.DrawImageOptions{}
			dio.GeoM.Translate(float64(i*2094), float64(j*1098))
			cam.Draw(assets.Rug, dio, screen)
		}
	}
}
