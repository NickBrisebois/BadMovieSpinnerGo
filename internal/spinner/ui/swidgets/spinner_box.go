package swidgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerBox struct {
	width     int
	height    int
	container *widget.Container
}

func NewSpinnerBox(bgColour color.Color) *SpinnerBox {
	spinnerRootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(bgColour),
		),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.DefaultStretch(true, true),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
	)
	return &SpinnerBox{
		container: spinnerRootContainer,
	}
}

func (g *SpinnerBox) GetContainer() *widget.Container {
	return g.container
}
