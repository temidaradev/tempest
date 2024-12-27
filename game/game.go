package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/setanarut/kamera/v2"
	"github.com/temidaradev/esset"
	"github.com/temidaradev/tempest/assets"
)

var (
	w, h             = 635., 475.
	groundY          = 400.
	unit             = 10.
	targetX, targetY = w / 2, h / 2
	cam              = kamera.NewCamera(targetX, targetY, w, h)
)

type Game struct {
	player *Player
	npc    *NPC
}

func init() {
	cam.SmoothType = kamera.SmoothDamp
	cam.SmoothOptions.SmoothDampTimeY = 0.2
}

func NewGame() *Game {
	g := &Game{}
	cam.ShakeEnabled = true

	g.player = &Player{
		DIO:   &ebiten.DrawImageOptions{},
		speed: 10,
	}

	g.npc = &NPC{
		DIO: &ebiten.DrawImageOptions{},
	}

	return g
}

var playerOffsetX = float64(assets.IdleTile[0].Bounds().Dx() / 2)
var playerOffsetY = float64(assets.IdleTile[0].Bounds().Dy() / 2)

func (g *Game) Draw(screen *ebiten.Image) {
	if g.player.isEntered[0] {
		HandleInterior(screen)
		g.npc.createNPC(screen)
	} else {
		HandleBackground(screen)
	}
	opO1 := &text.DrawOptions{}
	opO1.GeoM.Translate(g.player.X+playerOffsetX+325, g.player.Y+playerOffsetY+280)
	opO1.ColorScale.ScaleWithColor(color.White)

	opO2 := &text.DrawOptions{}
	opO2.GeoM.Translate(g.player.X+playerOffsetX+325, g.player.Y+playerOffsetY+296)
	opO2.ColorScale.ScaleWithColor(color.White)

	opI1 := &text.DrawOptions{}
	opI1.GeoM.Translate(g.player.X+playerOffsetX-200, g.player.Y+playerOffsetY+286)
	opI1.ColorScale.ScaleWithColor(color.White)

	opI2 := &text.DrawOptions{}
	opI2.GeoM.Translate(g.player.X+playerOffsetX-200, g.player.Y+playerOffsetY+300)
	opI2.ColorScale.ScaleWithColor(color.White)

	if g.player.isEntered[0] {
		if g.player.X >= 600 && g.player.X <= 670 {
			esset.UseFont(screen, assets.MyFont, "Go Back To Streets", 16, opI1)
			esset.UseFont(screen, assets.MyFont, "Press \"E\" to enter", 16, opI2)
		}
	} else {
		if g.player.X >= -150 && g.player.X <= -40 {
			esset.UseFont(screen, assets.MyFont, "Your Family's House", 16, opO1)
			esset.UseFont(screen, assets.MyFont, "Press \"E\" to enter", 16, opO2)
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %v Y:%v", g.player.X, g.player.Y))
	g.player.Draw(screen)
}

func (g *Game) Update() error {
	g.player.Update()
	g.npc.Update()
	if g.player.isEntered[0] {
		if g.player.X >= 0 && g.player.X <= 400 {
			cam.LookAt(g.player.X, g.player.Y)
		} else if g.player.X >= 400 {
			cam.LookAt(400, g.player.Y*2)
		} else if g.player.X <= 0 {
			cam.LookAt(0, g.player.Y*2)
		}
	} else {
		if g.player.X >= 0 && g.player.X <= 1660 {
			cam.LookAt(g.player.X, g.player.Y)
		} else if g.player.X >= 1660 {
			cam.LookAt(1660, g.player.Y*2)
		} else if g.player.X <= 0 {
			cam.LookAt(0, g.player.Y*2)
		}
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
