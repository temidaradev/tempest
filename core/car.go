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

func (c *Car) DrawCar(screen *ebiten.Image, color int, result int) {
	c.DIO.GeoM.Scale(2, 2)
	c.DIO.GeoM.Rotate(c.angle)
	c.DIO.GeoM.Translate(c.X+700, c.Y+450)

	msg := fmt.Sprintf("X: %+v, Y: %+v, Speed: %+v", c.X, c.Y, c.speed)
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
	turnSpeed := 0.01
	accel := 1.
	c.maxSpeed = 3.0
	c.friction = 0.1

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.speed += accel
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.speed -= accel
	}

	// No need to include speed/max speed here
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

	// calculate speed in x and y
	speedX := c.speed * math.Sin(c.angle)
	speedY := c.speed * math.Cos(c.angle)

	// update position based on velocity over time (this is meters per second * seconds = delta meters )
	c.X += speedX * 0.5
	c.Y += speedY * 0.5
}
