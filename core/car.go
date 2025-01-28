package game

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/temidaradev/tempest/assets"
)

type Car struct {
	X, Y     float64
	DIO      *ebiten.DrawImageOptions
	angle    float64
	speed    float64
	maxSpeed float64
	friction float64
}

func (c *Car) DrawCar(screen *ebiten.Image, color int) {
	c.DIO.GeoM.Scale(2, 2)
	c.DIO.GeoM.Rotate(c.angle)
	c.DIO.GeoM.Translate(c.X+700, c.Y+450)

	msg := fmt.Sprintf("X: %+v, Y: %+v, Speed: %+v", c.X, c.Y, c.speed)
	ebitenutil.DebugPrint(screen, msg)

	cam.Draw(assets.Compact[color-1], c.DIO, screen)
	c.DIO.GeoM.Reset()
}

func (c *Car) Update() {
	turnSpeed := 0.01
	accel := 1.
	c.maxSpeed = 3.0
	c.friction = 0.1

	// Hızlanma ve yavaşlama
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.speed += accel
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.speed -= accel
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.angle -= turnSpeed * (c.speed / c.maxSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.angle += turnSpeed * (c.speed / c.maxSpeed)
	}

	if c.speed > c.maxSpeed {
		c.speed = c.maxSpeed
	}
	if c.speed < -c.maxSpeed {
		c.speed = -c.maxSpeed
	}

	if c.speed > 0 {
		c.speed -= c.friction
		if c.speed < 0 {
			c.speed = 0
		}
	} else if c.speed < 0 {
		c.speed += c.friction
		if c.speed > 0 {
			c.speed = 0
		}
	}

	c.X += c.speed * math.Cos(c.angle)
	c.Y += c.speed * math.Sin(c.angle)
}
