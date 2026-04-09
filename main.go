package main

import (
	"log"
	"net/http"

	"NickBrisebois/BadMovieSpinnerGo/components"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	app.Route("/", func() app.Composer { return &components.HelloComponent{} })

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Bad Movie Spinner",
		Description: "A Bad Movie Spinner",
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
