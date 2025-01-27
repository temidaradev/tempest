package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/temidaradev/esset/v2"
	"github.com/temidaradev/tempest/assets"
)

type Menu struct {
	selection [2]bool
	kind      [4]bool
	index     int8
}

func (m *Menu) SelectCar(screen *ebiten.Image) {
	if m.selection[0] {
		esset.DrawText(screen, assets.MyFont, "Select A Color", 48, w/2-148, h/2-148, color.White)
		esset.DrawText(screen, assets.MyFont, "Press Enter To Select Color", 48, w/2-288, h/2-88, color.White)
		if m.index == 1 {
			for j, _ := range assets.Compact {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(j)*60, 500)
				screen.DrawImage(assets.Compact[j], op)
			}
		}
		if m.index == 2 {
			for i, _ := range assets.Coupe {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(i)*60, 500)
				screen.DrawImage(assets.Coupe[i], op)
			}
		}
		if m.index == 3 {
			for i, _ := range assets.Sedan {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(i)*60, 500)
				screen.DrawImage(assets.Sedan[i], op)
			}
		}
		if m.index == 4 {
			for i, _ := range assets.Sport {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(i)*60, 500)
				screen.DrawImage(assets.Sport[i], op)
			}
		}
	} else {
		esset.DrawText(screen, assets.MyFont, "Select A Car", 48, w/2-148, h/2-148, color.White)
		esset.DrawText(screen, assets.MyFont, "Press Arrow Keys To Navigate", 48, w/2-318, h/2-88, color.White)
	}
}

func (m *Menu) HandleSelectCar() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		m.index++
		fmt.Print(m.index)
		m.selection[0] = true
		//m.kind[m.index-1] = true
		if m.index >= 4 {
			m.index = 4
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		m.index--
		if m.index <= 1 {
			m.index = 1
		}
		fmt.Print(m.index)
		m.selection[0] = true
	}
	if m.selection[0] {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			fmt.Println("Color Selection")
		}
	}
	m.Button()
	return nil
}

func (m *Menu) Button() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.selection[0] = false
		m.kind[0] = false
		m.kind[1] = false
		m.kind[2] = false
		m.kind[3] = false
		m.index = 0
	}
}
