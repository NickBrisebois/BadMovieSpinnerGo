package spinnerbox

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"

	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerBox struct {
	container *widget.Container
	movies    *map[string][]models.MovieMeta
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

func (g *SpinnerBox) SetMovies(movies *map[string][]models.MovieMeta) {
	g.movies = movies
}
