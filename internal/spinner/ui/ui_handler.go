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
	sidebarWidgetWidthPercent = 15
)

type UIHandler struct {
	ui             *ebitenui.UI
	uiResources    *res.UIResources
	screenWidth    int
	screenHeight   int
	rootContainer  *widget.Container
	container      *widget.Container
	spinnerBox     *spinnerbox.SpinnerBox
	spinnerOverlay *spinnerbox.SpinnerOverlay
	sidebar        *sidebar.Sidebar
	logger         *slog.Logger
}

func NewUIHandler(screenWidth, screenHeight int, logger *slog.Logger) *UIHandler {
	// create the main content container
	container := widget.NewContainer(
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

	ui := ebitenui.UI{Container: container}

	uiResources, err := res.NewUIResources()
	if err != nil {
		logger.Error("failed to load UI resources", "error", err)
		return nil
	}

	handler := UIHandler{
		ui:           &ui,
		uiResources:  uiResources,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		container:    container,
		sidebar: sidebar.NewSidebar(
			screenWidth,
			sidebarWidgetWidthPercent,
			uiResources,
		),
		spinnerBox:     spinnerbox.NewSpinnerBox(uiResources),
		spinnerOverlay: spinnerbox.NewSpinnerOverlay(uiResources),
		logger:         logger,
	}
	container.AddChild(handler.sidebar.GetContainer())
	container.AddChild(handler.spinnerBox.GetContainer())

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
