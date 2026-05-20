package res

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources/assets"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func loadGraphic(path string, targetWidth, targetHeight int) (*ebiten.Image, error) {
	f, err := assets.Graphics.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	if targetWidth == 0 && targetHeight == 0 {
		return img, nil
	}

	// scale larger images to the target size if requested (there's probably a better way to do this?)
	scaleX := float64(targetWidth) / float64(img.Bounds().Dx())
	scaleY := float64(targetHeight) / float64(img.Bounds().Dy())

	scaledImg := ebiten.NewImage(targetWidth, targetHeight)
	scaledImgOpts := &ebiten.DrawImageOptions{}
	scaledImgOpts.GeoM.Scale(scaleX, scaleY)
	scaledImg.DrawImage(img, scaledImgOpts)

	return scaledImg, nil
}

type CheckboxResources struct {
	Image   *widget.CheckboxImage
	Spacing int
}

func loadCheckboxResources() (*CheckboxResources, error) {
	targetCheckboxWidth := 14
	targetCheckboxHeight := 14
	uncheckedImg, err := loadGraphic(imgCheckboxUnchecked, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	uncheckedHoveredImg, err := loadGraphic(imgCheckboxUncheckedHovered, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	uncheckedDisabledImg, err := loadGraphic(imgCheckboxUncheckedDisabled, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	checkedImg, err := loadGraphic(imgCheckboxChecked, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	checkedHoveredImg, err := loadGraphic(imgCheckboxCheckedHovered, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	checkedDisabledImg, err := loadGraphic(imgCheckboxCheckedDisabled, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	greyedImg, err := loadGraphic(imgCheckboxGreyed, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	greyedHoveredImg, err := loadGraphic(imgCheckboxGreyedHovered, targetCheckboxWidth, targetCheckboxHeight)
	if err != nil {
		return nil, err
	}
	greyedDisabledImg, err := loadGraphic(imgCheckboxGreyedDisabled, targetCheckboxWidth, targetCheckboxHeight)
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
