package swidgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Toolbar struct {
	container  *widget.Container
	Height     int
	fileMenu   *widget.Button
	helpButton *widget.Button
}

func NewToolbar(height int, bgColour color.Color) *Toolbar {
	toolbarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(bgColour),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{Stretch: true},
			),
			widget.WidgetOpts.MinSize(0, height),
		),
	)

	return &Toolbar{
		container: toolbarContainer,
	}
}

func (t *Toolbar) GetContainer() *widget.Container {
	return t.container
}
