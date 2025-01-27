package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/temidaradev/esset/v2"
	"github.com/temidaradev/tempest/assets"
)

type Menu struct {
	selection bool
	kind      [4]bool
}

func (m *Menu) SelectCar(screen *ebiten.Image) {
	if m.selection {
		esset.DrawText(screen, assets.MyFont, "Select A Color", 48, w/2-148, h/2-148, color.White)
		if m.kind[0] {
			for i, _ := range assets.Compact {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(i)*60, 400)
				screen.DrawImage(assets.Compact[i], op)
			}
		}
		if m.kind[1] {
			for i, _ := range assets.Coupe {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(i)*60, 400)
				screen.DrawImage(assets.Coupe[i], op)
			}
		}
		if m.kind[2] {
			for i, _ := range assets.Sedan {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(i)*60, 400)
				screen.DrawImage(assets.Sedan[i], op)
			}
		}
		if m.kind[3] {
			for i, _ := range assets.Sport {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(1.5, 1.5)
				op.GeoM.Translate(625+float64(i)*60, 400)
				screen.DrawImage(assets.Sport[i], op)
			}
		}
	} else {
		esset.DrawText(screen, assets.MyFont, "Select A Car \n1. Compact\n2. Coupe\n3. Sedan\n4. Sport", 48, w/2-148, h/2-148, color.White)
	}
}

func (m *Menu) HandleSelectCar() error {
	if !m.selection {
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			m.selection = true
			m.kind[0] = true
		}
		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			m.selection = true
			m.kind[1] = true
		}
		if inpututil.IsKeyJustPressed(ebiten.Key3) {
			m.selection = true
			m.kind[2] = true
		}
		if inpututil.IsKeyJustPressed(ebiten.Key4) {
			m.selection = true
			m.kind[3] = true
		}
	}
	m.Button()
	return nil
}

func (m *Menu) Button() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		m.selection = false
		m.kind[0] = false
		m.kind[1] = false
		m.kind[2] = false
		m.kind[3] = false
	}
}
