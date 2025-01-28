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
		if m.index == 1 {
			esset.DrawText(screen, assets.MyFont, "Compact", 48, w/2-138, h/2-48, color.White)
		}
		if m.index == 2 {
			esset.DrawText(screen, assets.MyFont, "Coupe", 48, w/2-138, h/2-48, color.White)
		}
		if m.index == 3 {
			esset.DrawText(screen, assets.MyFont, "Sedan", 48, w/2-138, h/2-48, color.White)
		}
		if m.index == 4 {
			esset.DrawText(screen, assets.MyFont, "Sport", 48, w/2-138, h/2-48, color.White)
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
			if m.index == 1 {
				fmt.Println("Compact")
			}
			if m.index == 2 {
				fmt.Println("Coupe")
			}
			if m.index == 3 {
				fmt.Println("Sedan")
			}
			if m.index == 4 {
				fmt.Println("Sport")
			}
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
