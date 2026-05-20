package res

import (
	_ "image"
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

	fontFaceRegularPath      = "fonts/NotoSans-Regular.ttf"
	fontFaceBoldPath         = "fonts/NotoSans-Bold.ttf"
	fontFaceIconsRegularPath = "fonts/FontAwesome7Free-Regular-400.otf"
	fontFaceIconsSolidPath   = "fonts/FontAwesome7Free-Solid-900.otf"

	imgCheckboxUnchecked         = "graphics/checkbox/checkbox.png"
	imgCheckboxUncheckedHovered  = "graphics/checkbox/checkbox.png"
	imgCheckboxUncheckedDisabled = "graphics/checkbox/checkbox.png"
	imgCheckboxChecked           = "graphics/checkbox/checkbox-checked.png"
	imgCheckboxCheckedHovered    = "graphics/checkbox/checkbox.png"
	imgCheckboxCheckedDisabled   = "graphics/checkbox/checkbox.png"
	imgCheckboxGreyed            = "graphics/checkbox/checkbox.png"
	imgCheckboxGreyedHovered     = "graphics/checkbox/checkbox.png"
	imgCheckboxGreyedDisabled    = "graphics/checkbox/checkbox.png"
	checkboxSpacing              = 10
)

var (
	ThemeBGColour                   = hexToColour(colourDuskBlue)
	ThemeSidebarBGColour            = hexToColour(colourDustyLavender)
	ThemeSidebarLabelIdleColour     = hexToColour(colourWhite)
	ThemeSidebarLabelDisabledColour = hexToColour(colourLightBronze)
	ThemeHeaderBGColour             = hexToColour(colourRosewood)
	ThemeBodyAccentColour           = hexToColour(colourRosewood)
	ThemeBodyTextColour             = hexToColour(colourWhite)
	ThemeHeaderTextColour           = hexToColour(colourWhite)
	ThemeHeaderAccentColour         = hexToColour(colourLightBronze)

	fontFaceRegular      = fontConfig{fontPath: fontFaceRegularPath, fontSize: 14}
	fontFaceBold         = fontConfig{fontPath: fontFaceBoldPath, fontSize: 14}
	fontFaceIconsRegular = fontConfig{fontPath: fontFaceIconsRegularPath, fontSize: 16}
	fontFaceIconsSolid   = fontConfig{fontPath: fontFaceIconsSolidPath, fontSize: 16}
)

type UIResources struct {
	Fonts          *FontResources
	Checkbox       *CheckboxResources
	LabelResources *LabelResources
}

func NewUIResources() (*UIResources, error) {
	checkbox, err := loadCheckboxResources()
	if err != nil {
		return nil, err
	}

	fontResources, err := loadFontResources(fontFaceRegular, fontFaceBold, fontFaceIconsRegular, fontFaceIconsSolid)
	if err != nil {
		return nil, err
	}

	labelResources := newLabelResources(
		labelStateColours{
			idle:     ThemeSidebarLabelIdleColour,
			disabled: ThemeSidebarLabelDisabledColour,
		},
		fontResources.FaceRegular,
	)

	return &UIResources{
		Fonts:          fontResources,
		LabelResources: labelResources,
		Checkbox:       checkbox,
	}, nil
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
