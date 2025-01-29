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

func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func (c *Car) Position() (int, int) {
	return int(c.X), int(c.Y)
}

func (c *Car) Angle() int {
	return int(c.angle)
}

func (c *Car) DrawCar(screen *ebiten.Image, color int, result int) {
	c.DIO.GeoM.Translate(-float64(32)/2, -float64(64)/2)
	c.DIO.GeoM.Rotate(DegreesToRadians(c.angle))
	c.DIO.GeoM.Scale(2, 2)
	c.DIO.GeoM.Translate(c.X+650, c.Y+350)

	msg := fmt.Sprintf("X: %+v, Y: %+v, FPS: %+v", c.X, c.Y, ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)

	if result == 1 {
		cam.Draw(assets.Compact[color-1], c.DIO, screen)
	}
	if result == 2 {
		cam.Draw(assets.Coupe[color-1], c.DIO, screen)
	}
	if result == 3 {
		cam.Draw(assets.Sedan[color-1], c.DIO, screen)
	}
	if result == 4 {
		cam.Draw(assets.Sport[color-1], c.DIO, screen)
	}
	c.DIO.GeoM.Reset()
}

func (c *Car) Update() {
	turnSpeed := 3.0
	accel := 1.
	c.maxSpeed = 7.0
	c.friction = 0.1

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.speed += accel * c.maxSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.speed -= accel * c.maxSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		c.maxSpeed = 12
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.angle -= turnSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.angle += turnSpeed
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

	c.X += c.speed * math.Sin(DegreesToRadians(c.angle))
	c.Y -= c.speed * math.Cos(DegreesToRadians(c.angle))

	if c.X >= 9780 {
		c.X = 9780
	}
	if c.Y <= -300 {
		c.Y = -300
	}
	if c.Y >= 5090 {
		c.Y = 5090
	}
	if c.X <= -605 {
		c.X = -605
	}
}
