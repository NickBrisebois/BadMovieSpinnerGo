package widgets

import "github.com/ebitenui/ebitenui/widget"

type Toolbar struct {
	container *widget.Container
}

func NewToolbar() *Toolbar {
	return &Toolbar{
		container: widget.NewContainer(),
	}
}

func (t *Toolbar) GetContainer() *widget.Container {
	return t.container
}
