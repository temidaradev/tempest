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
	clear     bool
	index     int8
	IndexC    int
	sDone     bool
}

func (m *Menu) SelectCar(screen *ebiten.Image) {
	if m.selection[0] {
		if m.index == 1 {
			esset.DrawText(screen, "Compact", 48, w/2-115, h/2-48, assets.FontFace, color.White)
			if m.selection[1] {
				assets.DrawCompactForMenu(screen)
				esset.DrawText(screen, "1   2   3   4", 48, w/2-125, h/2+125, assets.FontFace, color.White)
			}
		}
		if m.index == 2 {
			esset.DrawText(screen, "Coupe", 48, w/2-75, h/2-48, assets.FontFace, color.White)
			if m.selection[1] {
				assets.DrawCoupeForMenu(screen)
				esset.DrawText(screen, "1   2   3   4", 48, w/2-115, h/2+140, assets.FontFace, color.White)
			}
		}
		if m.index == 3 {
			esset.DrawText(screen, "Sedan", 48, w/2-75, h/2-48, assets.FontFace, color.White)
			if m.selection[1] {
				assets.DrawSedanForMenu(screen)
				esset.DrawText(screen, "1   2   3   4", 48, w/2-120, h/2+140, assets.FontFace, color.White)
			}
		}
		if m.index == 4 {
			esset.DrawText(screen, "Sport", 48, w/2-75, h/2-48, assets.FontFace, color.White)
			if m.selection[1] {
				assets.DrawSportForMenu(screen)
				esset.DrawText(screen, "1   2   3   4", 48, w/2-125, h/2+140, assets.FontFace, color.White)
			}
		}
	} else {
		esset.DrawText(screen, "Select A Car", 48, w/2-148, h/2-148, assets.FontFace, color.White)
		esset.DrawText(screen, "Press Arrow Keys To Navigate", 48, w/2-318, h/2-88, assets.FontFace, color.White)
	}
	if m.clear {
		screen.Clear()
	}
}

func (m *Menu) HandleSelectCar() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		if !m.selection[1] {
			m.index++
		}
		m.selection[0] = true
		if m.index >= 4 {
			m.index = 4
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		if !m.selection[1] {
			m.index--
		}
		if m.index <= 1 {
			m.index = 1
		}
		m.selection[0] = true
	}
	if m.selection[0] {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if m.index == 1 {
				fmt.Println("Compact")
				m.selection[1] = true
			}
			if m.index == 2 {
				fmt.Println("Coupe")
				m.selection[1] = true
			}
			if m.index == 3 {
				fmt.Println("Sedan")
				m.selection[1] = true
			}
			if m.index == 4 {
				fmt.Println("Sport")
				m.selection[1] = true
			}
		}
	}
	if m.selection[1] {
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			if m.index == 1 || m.index == 2 || m.index == 3 || m.index == 4 {
				m.IndexC = 1
				fmt.Println("Color 1")
				m.selection[1] = false
				m.clear = true
				m.sDone = true
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			if m.index == 1 || m.index == 2 || m.index == 3 || m.index == 4 {
				m.IndexC = 2
				fmt.Println("Color 2")
				m.selection[1] = false
				m.clear = true
				m.sDone = true
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.Key3) {
			if m.index == 1 || m.index == 2 || m.index == 3 || m.index == 4 {
				m.IndexC = 3
				fmt.Println("Color 3")
				m.selection[1] = false
				m.clear = true
				m.sDone = true
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.Key4) {
			if m.index == 1 || m.index == 2 || m.index == 3 || m.index == 4 {
				m.IndexC = 4
				fmt.Println("Color 4")
				m.selection[1] = false
				m.clear = true
				m.sDone = true
			}
		}
	}
	m.Button()
	return nil
}

func (m *Menu) Button() {
	if !m.sDone {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			m.selection[0] = false
			m.selection[1] = false
			m.index = 0
		}
	}
}
