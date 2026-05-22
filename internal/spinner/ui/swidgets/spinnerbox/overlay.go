package spinnerbox

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"

	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerOverlay struct {
	container   *widget.Container
	uiResources *res.UIResources
	movies      *map[string][]models.MovieMeta
}

func NewSpinnerOverlay(uiResources *res.UIResources) *SpinnerOverlay {
	// overlayRootContainer := widget.NewContainer(
	// 	// widget.ContainerOpts.Layout(
	// 	// )
	// )

	spinnerOverlay := &SpinnerOverlay{
		uiResources: uiResources,
		container:   widget.NewContainer(),
	}

	return spinnerOverlay
}
