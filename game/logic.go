package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Axis() (axisX, axisY float64) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {

	}
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		axisY = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		axisX = -0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		axisX = 0.5
	}
	return axisX, axisY
}
