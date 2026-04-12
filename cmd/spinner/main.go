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
	spinner    *spinner.Spinner
	isSpinning bool
}

func (g *BadMovieSpinner) Update() error {
	if g.isSpinning {
		g.spinner.Update()
	}
	return nil
}

func (g *BadMovieSpinner) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	g.spinner.DrawHandler.Draw(screen, screenWidth/2, screenHeight/2, screenWidth/2, screenHeight/2)
}

func (g *BadMovieSpinner) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	badMovieSpinner := BadMovieSpinner{
		spinner: spinner.NewSpinner(),
	}
	badMovieSpinner.spinner.Init(
		screenWidth/2,
		screenHeight/2,
		screenWidth/4,
		screenHeight/4,
	)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bad Movie Spinner")
	if err := ebiten.RunGame(&badMovieSpinner); err != nil {
		log.Fatal(err)
	}
}
