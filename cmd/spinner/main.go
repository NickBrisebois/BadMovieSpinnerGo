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
	spinner *spinner.DrawHandler
}

func (g *BadMovieSpinner) Update() error {
	return nil
}

func (g *BadMovieSpinner) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	g.spinner.Draw(screen, screenWidth/2, screenHeight/2, screenWidth/2, screenHeight/2)
}

func (g *BadMovieSpinner) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	spinner := spinner.NewDrawHandler(5)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bad Movie Spinner")
	if err := ebiten.RunGame(&BadMovieSpinner{spinner}); err != nil {
		log.Fatal(err)
	}
}
