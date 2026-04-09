package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type HelloComponent struct {
	app.Compo
}

func (h *HelloComponent) Render() app.UI {
	return app.H1().Text("BAD MOVIES test")
}
