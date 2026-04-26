package widgets

import "github.com/ebitenui/ebitenui/widget"

type Header struct {
	container *widget.Container
}

func NewHeader(title string) *Header {
	container := widget.NewContainer()
	return &Header{
		container: container,
	}
}
