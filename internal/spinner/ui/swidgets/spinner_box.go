package swidgets

import (
	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerBox struct {
	container *widget.Container
}

func NewSpinnerBox() *SpinnerBox {
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
