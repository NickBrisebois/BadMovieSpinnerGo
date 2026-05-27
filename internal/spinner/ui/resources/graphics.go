package res

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources/assets"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type scaleOpts struct {
	targetWidth  int
	targetHeight int
}

type graphicLoadOptions struct {
	// scale image being loaded before returning it to specified dimensions
	scaleOpts *scaleOpts
}

func loadGraphic(path string, loadOpts *graphicLoadOptions) (*ebiten.Image, error) {
	f, err := assets.Graphics.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := ebitenutil.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	if loadOpts == nil {
		return img, nil
	}

	if loadOpts.scaleOpts != nil {
		// scale larger images to the target size if requested (there's probably a better way to do this?)
		img = resizeImage(
			img,
			loadOpts.scaleOpts.targetWidth,
			loadOpts.scaleOpts.targetHeight,
		)
	}

	return img, nil
}

type CheckboxResources struct {
	Image   *widget.CheckboxImage
	Spacing int
}

func loadCheckboxResources() (*CheckboxResources, error) {
	targetCheckboxWidth := applyIntDeviceScale(14)
	targetCheckboxHeight := applyIntDeviceScale(14)
	scaleOpts := &scaleOpts{targetWidth: targetCheckboxWidth, targetHeight: targetCheckboxHeight}
	loadOpts := &graphicLoadOptions{scaleOpts: scaleOpts}

	uncheckedImg, err := loadGraphic(imgCheckboxUnchecked, loadOpts)
	if err != nil {
		return nil, err
	}
	uncheckedHoveredImg, err := loadGraphic(imgCheckboxUncheckedHovered, loadOpts)
	if err != nil {
		return nil, err
	}
	uncheckedDisabledImg, err := loadGraphic(imgCheckboxUncheckedDisabled, loadOpts)
	if err != nil {
		return nil, err
	}
	checkedImg, err := loadGraphic(imgCheckboxChecked, loadOpts)
	if err != nil {
		return nil, err
	}
	checkedHoveredImg, err := loadGraphic(imgCheckboxCheckedHovered, loadOpts)
	if err != nil {
		return nil, err
	}
	checkedDisabledImg, err := loadGraphic(imgCheckboxCheckedDisabled, loadOpts)
	if err != nil {
		return nil, err
	}
	greyedImg, err := loadGraphic(imgCheckboxGreyed, loadOpts)
	if err != nil {
		return nil, err
	}
	greyedHoveredImg, err := loadGraphic(imgCheckboxGreyedHovered, loadOpts)
	if err != nil {
		return nil, err
	}
	greyedDisabledImg, err := loadGraphic(imgCheckboxGreyedDisabled, loadOpts)
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

type ButtonResources struct {
	Image   *widget.ButtonImage
	Text    *widget.ButtonTextColor
	Face    *text.Face
	Padding *widget.Insets
}

func loadButtonResources(
	idleImgPath, hoverImgPath, pressedImgPath string,
	idleTextColour, hoverTextColour, pressedTextColour color.Color,
	fontFace *text.Face,
	padding *widget.Insets,
) (*ButtonResources, error) {
	loadOpts := &graphicLoadOptions{}
	loadOpts.scaleOpts = &scaleOpts{
		targetWidth:  150,
		targetHeight: 150,
	}
	buttonImgIdle, err := loadGraphic(idleImgPath, loadOpts)
	if err != nil {
		return nil, err
	}
	buttonImgHover, err := loadGraphic(hoverImgPath, loadOpts)
	if err != nil {
		return nil, err
	}
	buttonImgPressed, err := loadGraphic(pressedImgPath, loadOpts)
	if err != nil {
		return nil, err
	}

	button := &widget.ButtonImage{
		Idle:    image.NewFixedNineSlice(buttonImgIdle),
		Hover:   image.NewFixedNineSlice(buttonImgHover),
		Pressed: image.NewFixedNineSlice(buttonImgPressed),
	}

	return &ButtonResources{
		Image: button,
		Text: &widget.ButtonTextColor{
			Idle:     idleTextColour,
			Hover:    hoverTextColour,
			Disabled: pressedTextColour,
		},
		Face:    fontFace,
		Padding: padding,
	}, nil
}
