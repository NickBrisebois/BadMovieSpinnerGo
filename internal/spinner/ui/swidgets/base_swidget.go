package swidgets

// swidgets is the package for handlers managing the container states of spinner UI widgets
// named swidgets to avoid conflicts with ebitengine packages

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"

	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerWidget interface {
	GetContainer() *widget.Container
	SetMovies(movies *map[string][]models.MovieMeta)
}
