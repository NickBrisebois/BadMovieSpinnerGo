package spinnerbox

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/events"
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"

	"github.com/ebitenui/ebitenui/widget"
)

type SpinnerOverlay struct {
	container     *widget.Container
	uiResources   *res.UIResources
	movies        *map[string][]models.MovieMeta
	eventCallback events.EventCallback
}

func NewSpinnerOverlay(uiResources *res.UIResources, eventCallback events.EventCallback) *SpinnerOverlay {
	overlayRootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	spinnerOverlay := &SpinnerOverlay{
		uiResources:   uiResources,
		container:     overlayRootContainer,
		eventCallback: eventCallback,
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
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			s.eventCallback(&events.EventCallbackData{EventType: events.EventTypeSpinButtonClicked})
		}),
	)

	s.container.AddChild(button)
}
