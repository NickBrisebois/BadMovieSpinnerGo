package swidgets

// swidgets is the package for handlers managing the container states of spinner UI widgets
// named swidgets to avoid conflicts with ebitengine packages

import "github.com/ebitenui/ebitenui/widget"

type SWidgetHandler interface {
	GetContainer() *widget.Container
}

func calculatePercentOf(p int, total int) float64 {
	fPercent := float64(p)
	fTotal := float64(total)
	return fPercent / 100.0 * fTotal
}
