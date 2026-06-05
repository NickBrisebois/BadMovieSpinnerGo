package sidebar

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/events"
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"maps"
	"slices"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Sidebar struct {
	uiResources   *res.UIResources
	container     *widget.Container
	movies        *map[models.PersonName][]models.MovieMeta
	eventCallback events.EventCallback
}

func NewSidebar(width int, uiResources *res.UIResources, eventCallback events.EventCallback) *Sidebar {
	sidebarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(res.ThemeSidebarBGColour),
		),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(2),
				widget.GridLayoutOpts.DefaultStretch(true, true),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
			widget.WidgetOpts.MinSize(width, 0),
		),
	)

	return &Sidebar{
		container:     sidebarContainer,
		uiResources:   uiResources,
		eventCallback: eventCallback,
	}
}

func (s *Sidebar) GetContainer() *widget.Container {
	return s.container
}

func (s *Sidebar) SetMovies(movies *map[models.PersonName][]models.MovieMeta) {
	s.movies = movies
	suggestedByToggles := NewSuggestedByToggle(
		slices.Collect(maps.Keys(*movies)),
		s.uiResources,
		func(toggled []models.PersonName, args *widget.CheckboxChangedEventArgs) {
			s.eventCallback(&events.EventCallbackData{
				EventType:      events.EventTypeSuggestedByChanged,
				SuggestedUsers: &toggled,
			})
		},
	)
	s.container.AddChild(suggestedByToggles.GetContainer())
}
