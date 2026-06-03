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
	imgCheckboxUncheckedHovered  = "graphics/checkbox/checkbox-checked-hover.png"
	imgCheckboxUncheckedDisabled = "graphics/checkbox/checkbox.png"
	imgCheckboxChecked           = "graphics/checkbox/checkbox-checked.png"
	imgCheckboxCheckedHovered    = "graphics/checkbox/checkbox-checked.png"
	imgCheckboxCheckedDisabled   = "graphics/checkbox/checkbox-checked.png"
	imgCheckboxGreyed            = "graphics/checkbox/checkbox.png"
	imgCheckboxGreyedHovered     = "graphics/checkbox/checkbox.png"
	imgCheckboxGreyedDisabled    = "graphics/checkbox/checkbox.png"
	checkboxSpacing              = 5

	imgSpinButtonIdle    = "graphics/spin_button/unpressed.png"
	imgSpinButtonHover   = "graphics/spin_button/unpressed_hover.png"
	imgSpinButtonPressed = "graphics/spin_button/pressed.png"
)

var (
	ThemeBGColour = hexToColour(colourDuskBlue)

	ThemeSidebarBGColour            = hexToColour(colourDustyLavender)
	ThemeSidebarLabelIdleColour     = hexToColour(colourWhite)
	ThemeSidebarLabelDisabledColour = hexToColour(colourLightBronze)
	ThemeSidebarWidth               = 20

	ThemeHeaderBGColour = hexToColour(colourRosewood)

	ThemeBodyAccentColour = hexToColour(colourRosewood)
	ThemeBodyTextColour   = hexToColour(colourWhite)

	ThemeHeaderTextColour   = hexToColour(colourWhite)
	ThemeHeaderAccentColour = hexToColour(colourLightBronze)

	ThemeSpinButtonTextColourIdle    = hexToColour(colourWhite)
	ThemeSpinButtonTextColourHover   = hexToColour(colourLightBronze)
	ThemeSpinButtonTextColourPressed = hexToColour(colourDuskBlue)

	fontFaceRegular      = fontConfig{fontPath: fontFaceRegularPath, fontSize: 16}
	fontFaceBold         = fontConfig{fontPath: fontFaceBoldPath, fontSize: 16}
	fontFaceIconsRegular = fontConfig{fontPath: fontFaceIconsRegularPath, fontSize: 16}
	fontFaceIconsSolid   = fontConfig{fontPath: fontFaceIconsSolidPath, fontSize: 16}
)

type UIResources struct {
	Fonts               *FontResources
	Checkbox            *CheckboxResources
	LabelResources      *LabelResources
	SpinButtonResources *ButtonResources
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

	spinButtonResources, err := loadButtonResources(
		imgSpinButtonIdle, imgSpinButtonHover, imgSpinButtonPressed,
		ThemeSidebarLabelIdleColour, ThemeSidebarLabelIdleColour, ThemeSidebarLabelIdleColour,
		fontResources.FaceRegular,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &UIResources{
		Fonts:               fontResources,
		LabelResources:      labelResources,
		Checkbox:            checkbox,
		SpinButtonResources: spinButtonResources,
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
