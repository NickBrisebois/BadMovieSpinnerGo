package ui

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/widgets"
	"log/slog"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type UIHandler struct {
	ui            *ebitenui.UI
	rootContainer *widget.Container
	logger        *slog.Logger
}

func NewUIHandler(logger *slog.Logger) *UIHandler {
	// create the primary UI container
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true}),
			widget.GridLayoutOpts.Padding(&widget.Insets{
				Top:    20,
				Bottom: 20,
			}),
			widget.GridLayoutOpts.Spacing(0, 20),
		)),
	)
	ui := ebitenui.UI{Container: rootContainer}

	// use primary UI to init the handler to return
	handler := UIHandler{
		ui:            &ui,
		rootContainer: rootContainer,
		logger:        logger,
	}

	// but then init the various widgets
	handler.addSidebar()

	return &handler
}

func (u *UIHandler) addSidebar() {
	sidebar := widgets.NewSidebar(50, colornames.Gainsboro)
	u.rootContainer.AddChild(sidebar.GetContainer())
}

func (u *UIHandler) Update() error {
	u.ui.Update()
	return nil
}

func (u *UIHandler) Draw(screen *ebiten.Image) {
	u.ui.Draw(screen)
}
