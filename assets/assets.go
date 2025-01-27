package assets

import (
	"embed"
	_ "image/png"

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
