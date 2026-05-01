package res

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/assets"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	// Ocean Sunset palette: https://coolors.co/palette/355070-6d597a-b56576-e56b6f-eaac8b
	colourDuskBlue      = "355070"
	colourDustyLavender = "6d597a"
	colourRosewood      = "B56576"
	colourLightCoral    = "e56b6f"
	colourLightBronze   = "eaac8b"
	colourWhite         = "f4e9cd"

	fontFaceRegular      = "fonts/NotoSans-Regular.ttf"
	fontFaceBold         = "fonts/NotoSans-Bold.ttf"
	fontFaceIconsRegular = "fonts/FontAwesome7Free-Regular-400.otf"
	fontFaceIconsSolid   = "fonts/FontAwesome7Free-Solid-900.otf"
)

var (
	ThemeBGColour           = hexToColour(colourDuskBlue)
	ThemeSidebarBGColour    = hexToColour(colourDustyLavender)
	ThemeHeaderBGColour     = hexToColour(colourRosewood)
	ThemeBodyAccentColour   = hexToColour(colourRosewood)
	ThemeBodyTextColour     = hexToColour(colourWhite)
	ThemeHeaderTextColour   = hexToColour(colourWhite)
	ThemeHeaderAccentColour = hexToColour(colourLightBronze)

	ThemeFontFaceRegular      = loadFont(fontFaceRegular, 14)
	ThemeFontFaceBold         = loadFont(fontFaceBold, 16)
	ThemeFontFaceIconsRegular = loadFont(fontFaceIconsRegular, 16)
	ThemeFontFaceIconsSolid   = loadFont(fontFaceIconsSolid, 16)
)

func loadFont(fontPath string, fontSize float64) text.Face {
	fontFile, err := assets.Fonts.Open(fontPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	faceSource, err := text.NewGoTextFaceSource(fontFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &text.GoTextFace{
		Source: faceSource,
		Size:   fontSize,
	}
}

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
