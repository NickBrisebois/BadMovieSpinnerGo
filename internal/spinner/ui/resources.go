package ui

import (
	"image/color"
	"strconv"
)

const (
	// Ocean Sunset palette: https://coolors.co/palette/355070-6d597a-b56576-e56b6f-eaac8b
	colourDuskBlue      = "355070"
	colourDustyLavender = "6d597a"
	colourRosewood      = "B56576"
	colourLightCoral    = "e56b6f"
	colourLightBronze   = "eaac8b"
	colourWhite         = "f4e9cd"
)

var (
	ThemeBGColour           = hexToColour(colourDuskBlue)
	ThemeSidebarBGColour    = hexToColour(colourDustyLavender)
	ThemeHeaderBGColour     = hexToColour(colourRosewood)
	ThemeBodyAccentColour   = hexToColour(colourRosewood)
	ThemeBodyTextColour     = hexToColour(colourWhite)
	ThemeHeaderTextColour   = hexToColour(colourWhite)
	ThemeHeaderAccentColour = hexToColour(colourLightBronze)
)

func hexToColour(h string) color.Color {
	// Convert hex colour value to color struct
	// Borrowed from ebitenui demo:
	// https://github.com/ebitenui/ebitenui/blob/b1c31d852489cc8b4bc1d9581eaa75686875e5a7/_examples/demo/resources.go#L800-L812
	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		return color.RGBA{}
	}
	return color.RGBA{
		uint8(u & 0xff0000 >> 16),
		uint8(u & 0xff00 >> 8),
		uint8(u & 0xff),
		255,
	}
}
