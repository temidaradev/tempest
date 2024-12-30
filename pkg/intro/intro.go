package intro

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/tempest/assets"
)

type State struct {
	fadeIn  int
	fadeOut int
	w, h    int
}

func NewState() *State {
	return &State{
		fadeIn:  100,
		fadeOut: 100,
	}
}

func (s *State) Update() error {
	if s.fadeIn > 0 {
		s.fadeIn--
	} else if s.fadeOut > 0 {
		s.fadeOut--
	}
	return nil
}

func (s *State) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	if s.fadeIn > 0 {
		op.ColorScale.ScaleAlpha(1.0 - float32(s.fadeIn)/100)
	} else if s.fadeOut > 0 {
		op.ColorScale.ScaleAlpha(float32(s.fadeOut) / 100)
	} else {
		op.ColorScale.ScaleAlpha(0)
	}
	{
		ebi := assets.CityDay
		op.GeoM.Translate(float64(s.w)/2-float64(ebi.Bounds().Dy()/2), float64(s.h)/2-float64(ebi.Bounds().Dy())/2)
		screen.DrawImage(ebi, op)
	}
}

func (s *State) Layout(ow, oh int) (int, int) {
	s.w, s.h = ow, oh
	return ow, oh
}
