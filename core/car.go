package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/temidaradev/tempest/assets"
)

type Car struct {
	X, Y  float64
	DIO   *ebiten.DrawImageOptions
	speed float64
}

func (c *Car) DrawCar(screen *ebiten.Image, color int) {
	c.DIO.GeoM.Scale(2, 2)
	c.DIO.GeoM.Translate(c.X+700, c.Y+450)
	msg := fmt.Sprintf("X: %+v, Y: %+v", c.X, c.Y)
	ebitenutil.DebugPrint(screen, msg)
	cam.Draw(assets.Compact[color-1], c.DIO, screen)
	c.DIO.GeoM.Reset()
}

func (c *Car) Update() {
	x, y := Axis()
	c.X += x * c.speed
	c.Y += y * c.speed
}
