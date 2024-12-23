package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/temidaradev/esset"
	"github.com/temidaradev/tempest/assets"
)

type Player struct {
	X, Y      float64
	vx, vy    float64
	DIO       *ebiten.DrawImageOptions
	speed     float64
	count     int
	isEntered bool
}

func (p *Player) Update() {
	p.count++
	p.vx, p.vy = Axis()
	p.X += p.vx * p.speed
	p.Y += p.vy * p.speed

	//-150, -40
	p.X = min(max(p.X, -285), 1965)

	if p.X >= -150 && p.X <= -40 {
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			p.isEntered = true
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opSpecific := &text.DrawOptions{}
	opSpecific.GeoM.Translate(p.X+playerOffsetX+325, p.Y+playerOffsetY+280)
	opSpecific.ColorScale.ScaleWithColor(color.White)

	opSpecific2 := &text.DrawOptions{}
	opSpecific2.GeoM.Translate(p.X+playerOffsetX+325, p.Y+playerOffsetY+296)
	opSpecific2.ColorScale.ScaleWithColor(color.White)

	if p.vx != 0.5 && p.vx != -0.5 {
		frameCountIdle := 4
		i := (p.count / 8) % frameCountIdle

		p.DIO.GeoM.Scale(2, 2)
		p.DIO.GeoM.Translate(p.X+playerOffsetX+225, p.Y+playerOffsetY+315)

		cam.Draw(assets.IdleTile[i], p.DIO, screen)
		p.DIO.GeoM.Reset()
	}
	if p.vx == 0.5 {
		frameCountWalk := 6
		i := (p.count / 8) % frameCountWalk

		p.DIO.GeoM.Scale(2, 2)
		p.DIO.GeoM.Translate(p.X+playerOffsetX+225, p.Y+playerOffsetY+315)

		cam.Draw(assets.RunTile[i], p.DIO, screen)
		p.DIO.GeoM.Reset()
	}
	if p.vx == -0.5 {
		frameCountWalk := 6
		i := (p.count / 8) % -frameCountWalk

		p.DIO.GeoM.Scale(-2, 2)
		p.DIO.GeoM.Translate(p.X+playerOffsetX+325, p.Y+playerOffsetY+315)

		cam.Draw(assets.RunTile[i], p.DIO, screen)
		p.DIO.GeoM.Reset()
	}
	if p.X >= -150 && p.X <= -40 {
		esset.UseFont(screen, assets.MyFont, "Your Family's House", 16, opSpecific)
		esset.UseFont(screen, assets.MyFont, "Press \"E\" to enter", 16, opSpecific2)
	}
}
