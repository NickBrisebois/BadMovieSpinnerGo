package res

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources/assets"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadGraphic(path string) (*ebiten.Image, error) {
	f, err := assets.Graphics.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

type CheckboxResources struct {
	Image   *widget.CheckboxImage
	Spacing int
}

func loadCheckboxResources() (*CheckboxResources, error) {
	uncheckedImg, err := loadGraphic(imgCheckboxUnchecked)
	if err != nil {
		return nil, err
	}
	uncheckedHoveredImg, err := loadGraphic(imgCheckboxUncheckedHovered)
	if err != nil {
		return nil, err
	}
	uncheckedDisabledImg, err := loadGraphic(imgCheckboxUncheckedDisabled)
	if err != nil {
		return nil, err
	}
	checkedImg, err := loadGraphic(imgCheckboxChecked)
	if err != nil {
		return nil, err
	}
	checkedHoveredImg, err := loadGraphic(imgCheckboxCheckedHovered)
	if err != nil {
		return nil, err
	}
	checkedDisabledImg, err := loadGraphic(imgCheckboxCheckedDisabled)
	if err != nil {
		return nil, err
	}
	greyedImg, err := loadGraphic(imgCheckboxGreyed)
	if err != nil {
		return nil, err
	}
	greyedHoveredImg, err := loadGraphic(imgCheckboxGreyedHovered)
	if err != nil {
		return nil, err
	}
	greyedDisabledImg, err := loadGraphic(imgCheckboxGreyedDisabled)
	if err != nil {
		return nil, err
	}

	return &CheckboxResources{
		Image: &widget.CheckboxImage{
			Unchecked:         image.NewFixedNineSlice(uncheckedImg),
			UncheckedHovered:  image.NewFixedNineSlice(uncheckedHoveredImg),
			UncheckedDisabled: image.NewFixedNineSlice(uncheckedDisabledImg),
			Checked:           image.NewFixedNineSlice(checkedImg),
			CheckedHovered:    image.NewFixedNineSlice(checkedHoveredImg),
			CheckedDisabled:   image.NewFixedNineSlice(checkedDisabledImg),
			Greyed:            image.NewFixedNineSlice(greyedImg),
			GreyedHovered:     image.NewFixedNineSlice(greyedHoveredImg),
			GreyedDisabled:    image.NewFixedNineSlice(greyedDisabledImg),
		},
		Spacing: checkboxSpacing,
	}, nil
}
