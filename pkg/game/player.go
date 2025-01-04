package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/temidaradev/tempest/assets"
)

type Player struct {
	X, Y      float64
	vx, vy    float64
	DIO       *ebiten.DrawImageOptions
	speed     float64
	count     int
	isEntered [3]bool
}

func (p *Player) Update() {
	p.Y = 535
	p.count++
	p.vx, p.vy = Axis()
	p.X += p.vx * p.speed
	p.Y += p.vy * p.speed

	//-150, -40
	if p.isEntered[0] {
		p.X = min(max(p.X, -325), 670)
	} else {
		p.X = min(max(p.X, -500), 1675)
	}

	if p.X >= -150 && p.X <= -40 {
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			p.isEntered[0] = true
		}
	} else if p.X >= 1680 && p.X <= 1820 {
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			p.isEntered[1] = true
		}
	}

	if p.isEntered[0] {
		if p.X >= 600 && p.X <= 670 {
			if inpututil.IsKeyJustPressed(ebiten.KeyE) {
				p.isEntered[0] = false
				p.X = -30
			}
		}
	}

	if p.isEntered[0] {
		if p.X >= 0 && p.X <= 400 {
			cam.LookAt(p.X, p.Y)
		} else if p.X >= 400 {
			cam.LookAt(400, p.Y)
		} else if p.X <= 0 {
			cam.LookAt(0, p.Y)
		}
	} else {
		if p.X >= 0 && p.X <= 1700 {
			cam.LookAt(p.X, p.Y-60)
		} else if p.X >= 1700 {
			cam.LookAt(1700, p.Y-60)
		} else if p.X < 350 {
			cam.LookAt(0, p.Y-60)
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opSpecific := &text.DrawOptions{}
	opSpecific.GeoM.Translate(p.X+playerOffsetX+500, p.Y+playerOffsetY)
	opSpecific.ColorScale.ScaleWithColor(color.White)

	opSpecific2 := &text.DrawOptions{}
	opSpecific2.GeoM.Translate(p.X+playerOffsetX+500, p.Y+playerOffsetY)
	opSpecific2.ColorScale.ScaleWithColor(color.White)

	if p.vx != 0.5 && p.vx != -0.5 {
		frameCountIdle := 4
		i := (p.count / 8) % frameCountIdle

		if p.isEntered[0] {
			p.DIO.GeoM.Scale(5, 5)
			p.DIO.GeoM.Translate(p.X+playerOffsetX+225, p.Y+playerOffsetY+600)

			cam.Draw(assets.IdleTile[i], p.DIO, screen)
		} else {
			p.DIO.GeoM.Scale(2, 2)
			p.DIO.GeoM.Translate(p.X+playerOffsetX+500, p.Y+playerOffsetY+600)

			cam.Draw(assets.IdleTile[i], p.DIO, screen)
		}
		p.DIO.GeoM.Reset()
	}
	if p.vx == 0.5 {
		frameCountWalk := 6
		i := (p.count / 8) % frameCountWalk

		if p.isEntered[0] {
			p.DIO.GeoM.Scale(5, 5)
			p.DIO.GeoM.Translate(p.X+playerOffsetX+225, p.Y+playerOffsetY+600)

			cam.Draw(assets.RunTile[i], p.DIO, screen)
		} else {
			p.DIO.GeoM.Scale(2, 2)
			p.DIO.GeoM.Translate(p.X+playerOffsetX+500, p.Y+playerOffsetY+600)

			cam.Draw(assets.RunTile[i], p.DIO, screen)
		}
		p.DIO.GeoM.Reset()
	}
	if p.vx == -0.5 {
		frameCountWalk := 6
		i := (p.count / 8) % -frameCountWalk

		if p.isEntered[0] {
			p.DIO.GeoM.Scale(-5, 5)
			p.DIO.GeoM.Translate(p.X+playerOffsetX+415, p.Y+playerOffsetY+600)

			cam.Draw(assets.RunTile[i], p.DIO, screen)
		} else {
			p.DIO.GeoM.Scale(-2, 2)
			p.DIO.GeoM.Translate(p.X+playerOffsetX+600, p.Y+playerOffsetY+600)

			cam.Draw(assets.RunTile[i], p.DIO, screen)
		}
		p.DIO.GeoM.Reset()
	}
}
