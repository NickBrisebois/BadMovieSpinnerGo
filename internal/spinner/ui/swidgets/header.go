package swidgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Header struct {
	container *widget.Container
}

func NewHeader(title string, bgColour color.Color) *Header {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(15),
		)),
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(bgColour),
		),
	)
	return &Header{
		container: container,
	}
}

func (h *Header) GetContainer() *widget.Container {
	return h.container
}
