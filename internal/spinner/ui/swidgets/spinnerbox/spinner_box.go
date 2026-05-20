package spinnerbox

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"

	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerBox struct {
	container   *widget.Container
	uiResources *res.UIResources
	movies      *map[string][]models.MovieMeta
}

func NewSpinnerBox(uiResources *res.UIResources) *SpinnerBox {
	spinnerRootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.DefaultStretch(true, true),
			),
		),
	)
	return &SpinnerBox{
		container:   spinnerRootContainer,
		uiResources: uiResources,
	}
}

func (g *SpinnerBox) GetContainer() *widget.Container {
	return g.container
}

func (g *SpinnerBox) SetMovies(movies *map[string][]models.MovieMeta) {
	g.movies = movies
}
