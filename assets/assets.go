package assets

import (
	"embed"

	"github.com/temidaradev/esset"
)

//go:embed *
var Assets embed.FS

var RunTile = esset.GetMultiAssets(Assets, "walk/*.png")
var IdleTile = esset.GetMultiAssets(Assets, "idle/*.png")
var CityNight = esset.GetAsset(Assets, "backgrounds/night.png")
var CityDay = esset.GetAsset(Assets, "backgrounds/day.png")
