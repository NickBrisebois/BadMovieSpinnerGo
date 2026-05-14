package ui

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/swidgets"
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
	// contentContainer holds the main content under the toolbar
	contentContainer *widget.Container
	spinnerBox       *swidgets.SpinnerBox
	sidebar          *swidgets.Sidebar
	toolbar          *swidgets.Header
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
				[]bool{false, true},
			),
		)),
	)

	// header gets its own fancy container above the main content
	header := swidgets.NewHeader(screenWidth, headerWidgetHeightPercent, res.ThemeHeaderBGColour)

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

	rootContainer.AddChild(header.GetContainer(), contentContainer)
	ui := ebitenui.UI{Container: rootContainer}

	// use primary UI to init the handler to return
	handler := UIHandler{
		ui:               &ui,
		screenWidth:      screenWidth,
		screenHeight:     screenHeight,
		rootContainer:    rootContainer,
		contentContainer: contentContainer,
		toolbar:          header,
		sidebar:          swidgets.NewSidebar(screenWidth, sidebarWidgetWidthPercent, res.ThemeSidebarBGColour),
		spinnerBox:       swidgets.NewSpinnerBox(),
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

func (u *UIHandler) Update() error {
	u.ui.Update()
	return nil
}

func (u *UIHandler) DrawUI(screen *ebiten.Image) {
	u.ui.Draw(screen)
}

func (u *UIHandler) DrawSpinnerOverlay(screen *ebiten.Image) {

}
