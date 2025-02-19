package assets

import (
	"embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/temidaradev/esset/v2"
)

//go:embed *
var Assets embed.FS

//go:embed fonts/Roboto-Bold.ttf
var MyFont []byte
var FontFace text.Face

// Splash
var Splash = esset.GetAsset(Assets, "splash/splash.jpg")
var SplashFull = esset.GetAsset(Assets, "splash/splash_1920x1080_black.png")

// Rug
var Rug = esset.GetAsset(Assets, "background/rug.png")

// Cars
var Compact = esset.GetMultiAssets(Assets, "cars/Compact/*.png")
var Coupe = esset.GetMultiAssets(Assets, "cars/Coupe/*.png")
var Sedan = esset.GetMultiAssets(Assets, "cars/Sedan/*.png")
var Sport = esset.GetMultiAssets(Assets, "cars/Sport/*.png")

func DrawCompactForMenu(screen *ebiten.Image) {
	for j, _ := range Compact {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(835+float64(j)*60, 560)
		screen.DrawImage(Compact[j], op)
	}
}

func DrawCoupeForMenu(screen *ebiten.Image) {
	for j, _ := range Coupe {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(845+float64(j)*60, 560)
		screen.DrawImage(Coupe[j], op)
	}
}

func DrawSedanForMenu(screen *ebiten.Image) {
	for j, _ := range Sedan {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(840+float64(j)*60, 560)
		screen.DrawImage(Sedan[j], op)
	}
}

func DrawSportForMenu(screen *ebiten.Image) {
	for j, _ := range Sport {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.5, 1.5)
		op.GeoM.Translate(835+float64(j)*60, 560)
		screen.DrawImage(Sport[j], op)
	}
}
