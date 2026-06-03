package sidebar

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"maps"
	"slices"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type SidebarInputType int

const (
	SidebarInputTypeSuggestedBy SidebarInputType = iota
)

type SidebarInputCallbackData struct {
	InputType      SidebarInputType
	SuggestedUsers *[]string
}

type SidebarInputCallback func(data *SidebarInputCallbackData)

type Sidebar struct {
	uiResources   *res.UIResources
	container     *widget.Container
	movies        *map[string][]models.MovieMeta
	inputCallback SidebarInputCallback
}

func NewSidebar(width int, uiResources *res.UIResources, inputCallback SidebarInputCallback) *Sidebar {
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
		inputCallback: inputCallback,
	}
}

func (s *Sidebar) GetContainer() *widget.Container {
	return s.container
}

func (s *Sidebar) toggleListCallback(toggled []string, args *widget.CheckboxChangedEventArgs) {
}

func (s *Sidebar) SetMovies(movies *map[string][]models.MovieMeta) {
	s.movies = movies
	suggestedByToggles := NewSuggestedByToggle(
		slices.Collect(maps.Keys(*movies)),
		s.uiResources,
		func(toggled []string, args *widget.CheckboxChangedEventArgs) {
			s.inputCallback(&SidebarInputCallbackData{
				InputType:      SidebarInputTypeSuggestedBy,
				SuggestedUsers: &toggled,
			})
		},
	)
	s.container.AddChild(suggestedByToggles.GetContainer())
}
