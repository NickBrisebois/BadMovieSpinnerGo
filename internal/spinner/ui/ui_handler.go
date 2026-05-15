package ui

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/swidgets/sidebar"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/swidgets/spinnerbox"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"image"
	"log/slog"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	headerWidgetHeightPercent = 3
	sidebarWidgetWidthPercent = 15
)

type UIHandler struct {
	ui           *ebitenui.UI
	screenWidth  int
	screenHeight int
	// rootContainer contains *everything*
	rootContainer *widget.Container
	// contentContainer holds the main content under the toolbar* (*toolbar currently un-implemented)
	contentContainer *widget.Container
	spinnerBox       *spinnerbox.SpinnerBox
	sidebar          *sidebar.Sidebar
	logger           *slog.Logger
}

func NewUIHandler(screenWidth, screenHeight int, logger *slog.Logger) *UIHandler {
	// all widgets live inside the root container
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Spacing(0, 0),
			widget.GridLayoutOpts.Stretch(
				[]bool{true},
				[]bool{true},
			),
		)),
	)

	// create the main content container
	contentContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Spacing(0, 0),
			widget.GridLayoutOpts.Stretch(
				[]bool{false, true},
				[]bool{true},
			),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(screenWidth, 0),
		),
	)

	ui := ebitenui.UI{Container: rootContainer}
	rootContainer.AddChild(contentContainer)

	// use primary UI to init the handler to return
	handler := UIHandler{
		ui:               &ui,
		screenWidth:      screenWidth,
		screenHeight:     screenHeight,
		rootContainer:    rootContainer,
		contentContainer: contentContainer,
		sidebar:          sidebar.NewSidebar(screenWidth, sidebarWidgetWidthPercent, res.ThemeSidebarBGColour),
		spinnerBox:       spinnerbox.NewSpinnerBox(),
		logger:           logger,
	}
	contentContainer.AddChild(handler.sidebar.GetContainer())
	contentContainer.AddChild(handler.spinnerBox.GetContainer())

	return &handler
}

func (u *UIHandler) GetSpinnerBoxRect() image.Rectangle {
	return u.spinnerBox.GetContainer().GetWidget().Rect
}

func (u *UIHandler) SetDimensions(screenWidth, screenHeight int) {
	u.screenWidth = screenWidth
	u.screenHeight = screenHeight
}

func (u *UIHandler) SetMovies(movies *map[string][]models.MovieMeta) {
	u.sidebar.SetMovies(movies)
}

func (u *UIHandler) Update() error {
	u.ui.Update()
	return nil
}

func (u *UIHandler) DrawUI(screen *ebiten.Image) {
	u.ui.Draw(screen)
}

func (u *UIHandler) DrawSpinnerOverlay(screen *ebiten.Image) {

}
