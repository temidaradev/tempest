package assets

import (
	"embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/esset/v2"
)

//go:embed *
var Assets embed.FS

//go:embed fonts/Roboto-Bold.ttf
var MyFont []byte

// Splash
var Splash = esset.GetAsset(Assets, "splash/splash.jpg")

// Cars
var Compact = esset.GetMultiAssets(Assets, "cars/Compact/*.png")
var Coupe = esset.GetMultiAssets(Assets, "cars/Coupe/*.png")
var Sedan = esset.GetMultiAssets(Assets, "cars/Sedan/*.png")
var Sport = esset.GetMultiAssets(Assets, "cars/Sport/*.png")

func DrawCompactForMenu(screen *ebiten.Image) {
	for j, _ := range Compact {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(625+float64(j)*60, 500)
		screen.DrawImage(Compact[j], op)
	}
}

func CompactGetColor(screen *ebiten.Image, index int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	screen.DrawImage(Compact[index], op)
}

func DrawCoupeForMenu(screen *ebiten.Image) {
	for j, _ := range Compact {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(625+float64(j)*60, 500)
		screen.DrawImage(Coupe[j], op)
	}
}

func CoupetGetColor(screen *ebiten.Image, index int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	screen.DrawImage(Coupe[index], op)
}

func DrawSedanForMenu(screen *ebiten.Image) {
	for j, _ := range Compact {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(625+float64(j)*60, 500)
		screen.DrawImage(Sedan[j], op)
	}
}

func SedanGetColor(screen *ebiten.Image, index int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	screen.DrawImage(Sedan[index], op)
}

func DrawSportForMenu(screen *ebiten.Image) {
	for j, _ := range Compact {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(625+float64(j)*60, 500)
		screen.DrawImage(Sport[j], op)
	}
}

func SportGetColor(screen *ebiten.Image, index int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	screen.DrawImage(Sport[index], op)
}
