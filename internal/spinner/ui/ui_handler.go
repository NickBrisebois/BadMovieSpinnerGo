package ui

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/swidgets"
	"image"
	"log/slog"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
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
	toolbar          *swidgets.Toolbar
	logger           *slog.Logger
}

func NewUIHandler(screenWidth, screenHeight int, logger *slog.Logger) *UIHandler {
	// all widgets live inside the root container
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(0),
		)),
	)

	// toolbar gets its own fancy container above the main content
	toolbar := swidgets.NewToolbar(35, colornames.Whitesmoke)

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
			widget.WidgetOpts.MinSize(screenWidth, screenHeight-toolbar.Height),
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
	)

	rootContainer.AddChild(toolbar.GetContainer(), contentContainer)
	ui := ebitenui.UI{Container: rootContainer}

	// use primary UI to init the handler to return
	handler := UIHandler{
		ui:               &ui,
		screenWidth:      screenWidth,
		screenHeight:     screenHeight,
		rootContainer:    rootContainer,
		contentContainer: contentContainer,
		toolbar:          toolbar,
		sidebar:          swidgets.NewSidebar(screenWidth, 20, colornames.Aquamarine),
		spinnerBox:       swidgets.NewSpinnerBox(colornames.Blanchedalmond),
		logger:           logger,
	}
	contentContainer.AddChild(handler.sidebar.GetContainer())
	contentContainer.AddChild(handler.spinnerBox.GetContainer())

	return &handler
}

func (u *UIHandler) GetSpinnerBoxRect() image.Rectangle {
	return u.spinnerBox.GetContainer().GetWidget().Rect
}

func (u *UIHandler) Update() error {
	u.ui.Update()
	return nil
}

func (u *UIHandler) Draw(screen *ebiten.Image) {
	u.ui.Draw(screen)
}
