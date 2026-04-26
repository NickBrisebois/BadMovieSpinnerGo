package widgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Sidebar struct {
	width     int
	container *widget.Container
}

func NewSidebar(width int, backgroundColour color.Color) *Sidebar {
	colourOpts := widget.ContainerOpts.BackgroundImage(
		image.NewNineSliceColor(backgroundColour),
	)
	widgetOpts := widget.ContainerOpts.WidgetOpts(
		widget.WidgetOpts.MinSize(width, 0),
		widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Stretch:  true,
				Position: widget.RowLayoutPositionEnd,
			},
		),
	)

	rootContainer := widget.NewContainer(colourOpts, widgetOpts)
	return &Sidebar{
		width:     width,
		container: rootContainer,
	}
}

func (s *Sidebar) GetContainer() *widget.Container {
	return s.container
}
