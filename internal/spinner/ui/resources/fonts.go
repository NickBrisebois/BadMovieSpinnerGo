package res

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources/assets"
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type FontResources struct {
	FaceRegular      *text.Face
	FaceBold         *text.Face
	FaceIconsRegular *text.Face
	FaceIconsSolid   *text.Face
}

type fontConfig struct {
	fontPath string
	fontSize float64
}

func loadFont(fontPath string, fontSize float64) (text.Face, error) {
	fontFile, err := assets.Fonts.Open(fontPath)
	if err != nil {
		return nil, err
	}
	defer fontFile.Close()

	faceSource, err := text.NewGoTextFaceSource(fontFile)
	if err != nil {
		return nil, err
	}

	return &text.GoTextFace{
		Source: faceSource,
		Size:   fontSize,
	}, nil
}

func loadFontResources(
	faceRegular, faceBold, faceIconsRegular, faceIconsSolid fontConfig,
) (*FontResources, error) {
	regular, err := loadFont(faceRegular.fontPath, faceRegular.fontSize)
	if err != nil {
		return nil, err
	}

	bold, err := loadFont(faceBold.fontPath, faceBold.fontSize)
	if err != nil {
		return nil, err
	}

	iconsRegular, err := loadFont(faceIconsRegular.fontPath, faceIconsRegular.fontSize)
	if err != nil {
		return nil, err
	}

	iconsSolid, err := loadFont(faceIconsSolid.fontPath, faceIconsSolid.fontSize)
	if err != nil {
		return nil, err
	}

	return &FontResources{
		FaceRegular:      &regular,
		FaceBold:         &bold,
		FaceIconsRegular: &iconsRegular,
		FaceIconsSolid:   &iconsSolid,
	}, nil
}

type LabelResources struct {
	Text *widget.LabelColor
	Face *text.Face
}

type labelStateColours struct {
	idle     color.Color
	disabled color.Color
}

func newLabelResources(stateColours labelStateColours, font *text.Face) *LabelResources {
	return &LabelResources{
		Text: &widget.LabelColor{
			Idle:     stateColours.idle,
			Disabled: stateColours.disabled,
		},
		Face: font,
	}
}
