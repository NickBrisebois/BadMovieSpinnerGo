package main

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 600
	screenHeight = 600
)

type BadMovieSpinner struct {
	spinner *spinner.Spinner
}

func (g *BadMovieSpinner) Update() error {
	g.spinner.Update()
	return nil
}

func (g *BadMovieSpinner) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	g.spinner.DrawHandler.Draw(screen)
}

func (g *BadMovieSpinner) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	badMovieSpinner := BadMovieSpinner{spinner: &spinner.Spinner{}}
	badMovieSpinner.spinner.Init(
		screenWidth/2,
		screenHeight/2,
		screenWidth/4,
		screenHeight/4,
	)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bad Movie Spinner")
	if err := ebiten.RunGame(&badMovieSpinner); err != nil {
		log.Fatal(err)
	}
}
