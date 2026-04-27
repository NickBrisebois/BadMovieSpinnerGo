package swidgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerBox struct {
	container *widget.Container
}

func NewSpinnerBox(bgColour color.Color) *SpinnerBox {
	spinnerRootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.DefaultStretch(true, true),
			),
		),
	)
	return &SpinnerBox{
		container: spinnerRootContainer,
	}
}

func (g *SpinnerBox) GetContainer() *widget.Container {
	return g.container
}
