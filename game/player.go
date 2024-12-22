package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/temidaradev/ehh24/assets"
)

type Player struct {
	X, Y   float64
	vx, vy float64
	DIO    *ebiten.DrawImageOptions
	speed  float64
	count  int
}

func (p *Player) Update() {
	p.count++
	p.vx, p.vy = Axis()
	p.X += p.vx * p.speed
	p.Y += p.vy * p.speed

	//-150, -40
	p.X = min(max(p.X, -285), 1965)
}

func (p *Player) Draw(screen *ebiten.Image) {
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
	if p.X >= -150 && p.X <= -45 {
		vector.DrawFilledRect(screen, float32(p.X+playerOffsetX+325), float32(p.Y+playerOffsetY+280), 140, 30, color.RGBA{155, 0, 255, 255}, true)
		ebitenutil.DebugPrintAt(screen, "Your Family's House\nPress \"E\" to enter", int(p.X+playerOffsetX+325), int(p.Y+playerOffsetY+280))
	}
}
