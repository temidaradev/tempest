package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/setanarut/kamera/v2"
	"github.com/temidaradev/esset/v2"
	"github.com/temidaradev/tempest/assets"
	"github.com/temidaradev/tempest/pkg/intro"
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
	intro  *intro.State
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

	g.intro = &intro.State{}

	return g
}

var playerOffsetX = float64(assets.IdleTile[0].Bounds().Dx() / 2)
var playerOffsetY = float64(assets.IdleTile[0].Bounds().Dy() / 2)

func (g *Game) Draw(screen *ebiten.Image) {
	g.intro.Draw(screen)
	if g.player.isEntered[0] {
		HandleInterior(screen)
		g.npc.createNPC(screen)
	} else {
		HandleBackground(screen)
	}

	if g.player.isEntered[0] {
		if g.player.X >= 600 && g.player.X <= 670 {
			esset.DrawText(screen, assets.MyFont, "Go Back To Streets\nPress \"E\" to enter", 16, g.player.X+playerOffsetX-200, g.player.Y+playerOffsetY+275, color.White)
		}
	} else {
		if g.player.X >= -150 && g.player.X <= -40 {
			esset.DrawText(screen, assets.MyFont, "Your Family's House\nPress \"E\" to enter", 16, g.player.X+playerOffsetX+200, g.player.Y+playerOffsetY+325, color.White)
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %v Y:%v", g.player.X, g.player.Y))
	g.player.Draw(screen)
}

func (g *Game) Update() error {
	g.intro.Update()
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
