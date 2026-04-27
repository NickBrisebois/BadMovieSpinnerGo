package ui

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/swidgets"
	"log/slog"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type UIHandler struct {
	ui               *ebitenui.UI
	screenWidth      int
	screenHeight     int
	rootContainer    *widget.Container
	contentContainer *widget.Container
	toolbar          *swidgets.Toolbar
	widgets          []swidgets.SWidgetHandler
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
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(screenWidth, screenHeight-toolbar.Height),
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
				StretchHorizontal:  true,
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
		logger:           logger,
	}

	sidebar := swidgets.NewSidebar(screenWidth, 20, colornames.Aquamarine)
	spinnerBox := swidgets.NewSpinnerBox(colornames.Blanchedalmond)

	// but then init the various widgets
	handler.addUIWidgets(sidebar, spinnerBox)
	return &handler
}

func (u *UIHandler) addUIWidgets(newWidgets ...swidgets.SWidgetHandler) {
	for _, w := range newWidgets {
		u.contentContainer.AddChild(w.GetContainer())
		u.widgets = append(u.widgets, w)
	}
}

func (u *UIHandler) GetSpinnerPosition(screenWidth, screenHeight int) float32 {
	return 0
}

func (u *UIHandler) Update() error {
	u.ui.Update()
	return nil
}

func (u *UIHandler) Draw(screen *ebiten.Image) {
	u.ui.Draw(screen)
}
