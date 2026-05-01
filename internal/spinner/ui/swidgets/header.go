package swidgets

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Header struct {
	container     *widget.Container
	HeightPercent int
	fileMenu      *widget.Button
	helpButton    *widget.Button
}

func NewHeader(screenHeight, heightPercent int, bgColour color.Color) *Header {
	headerHeight := calculatePercentOf(screenHeight, heightPercent)
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
			widget.WidgetOpts.MinSize(0, int(headerHeight)),
		),
	)

	toolbarContainer.AddChild(
		widget.NewText(
			widget.TextOpts.Text("Bad Movie Spinner", &res.ThemeFontFaceBold, res.ThemeHeaderTextColour),
		),
	)

	return &Header{
		container: toolbarContainer,
	}
}

func (t *Header) GetContainer() *widget.Container {
	return t.container
}
