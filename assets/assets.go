package assets

import (
	"embed"

	"github.com/temidaradev/esset"
)

//go:embed *
var Assets embed.FS

//go:embed fonts/Roboto-Bold.ttf
var MyFont []byte

// MC
var RunTile = esset.GetMultiAssets(Assets, "walk/*.png")
var IdleTile = esset.GetMultiAssets(Assets, "idle/*.png")

// BG
var CityNight = esset.GetAsset(Assets, "backgrounds/night.png")
var CityDay = esset.GetAsset(Assets, "backgrounds/day.png")
var Interior = esset.GetAsset(Assets, "backgrounds/interior.png")

// NPC
var DrunkTile = esset.GetMultiAssets(Assets, "npc/drunk/walk/*.png")
var ExBfTile = esset.GetMultiAssets(Assets, "npc/ex-bf/idle/*.png")
