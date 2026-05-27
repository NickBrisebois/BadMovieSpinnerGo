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
	overlayRootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	spinnerOverlay := &SpinnerOverlay{
		uiResources: uiResources,
		container:   overlayRootContainer,
	}

	spinnerOverlay.addSpinButton()

	return spinnerOverlay
}

func (s *SpinnerOverlay) GetContainer() *widget.Container {
	return s.container
}

func (s *SpinnerOverlay) addSpinButton() {
	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ButtonOpts.Image(s.uiResources.SpinButtonResources.Image),
		widget.ButtonOpts.Text("Spin", s.uiResources.SpinButtonResources.Face, s.uiResources.SpinButtonResources.Text),
		widget.ButtonOpts.TextPadding(s.uiResources.SpinButtonResources.Padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {}),
	)

	s.container.AddChild(button)
}
