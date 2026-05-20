package sidebar

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	swidgetutils "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/swidgets/utils"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"maps"
	"slices"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type Sidebar struct {
	screenWidth  int
	widthPercent int // width as a percentage of the screen width (eg. a value of 50 == 50% of screenWidth)
	uiResources  *res.UIResources
	container    *widget.Container
	movies       *map[string][]models.MovieMeta
}

func NewSidebar(screenWidth int, widthPercent int, uiResources *res.UIResources) *Sidebar {
	sidebarWidth := swidgetutils.CalculatePercentOf(screenWidth, widthPercent)
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
			widget.WidgetOpts.MinSize(int(sidebarWidth), 0),
		),
	)

	return &Sidebar{
		screenWidth:  screenWidth,
		widthPercent: widthPercent,
		container:    sidebarContainer,
		uiResources:  uiResources,
	}
}

func (s *Sidebar) GetSidebarWidth() int {
	return int(swidgetutils.CalculatePercentOf(s.screenWidth, s.widthPercent))
}

func (s *Sidebar) GetContainer() *widget.Container {
	return s.container
}

func (s *Sidebar) SetMovies(movies *map[string][]models.MovieMeta) {
	s.movies = movies
	suggestedByToggles := NewSuggestedByToggle(
		slices.Collect(maps.Keys(*movies)),
		s.uiResources,
	)
	s.container.AddChild(suggestedByToggles.GetContainer())
}
