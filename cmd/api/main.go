package main

import (
	"log"
	"reflect"

	"NickBrisebois/BadMovieSpinnerGo/internal/api"

	_ "NickBrisebois/BadMovieSpinnerGo/docs"
)

type SpinnerServerConfig struct {
	TMDBAPIKey     string `env:"TMDB_API_KEY"`
	GoogleScopes   string `env:"GOOGLE_SCOPES"`
	GoogleSheetsID string `env:"GOOGLE_SHEETS_ID"`
}

func (c *SpinnerServerConfig) Load() {

	t := reflect.TypeFor[SpinnerServerConfig]()

	for field := range t.Fields() {
		log.Println(field)
		// envName := field.Tag.Get("env")
		// if envName != "" {
		// 	envVal := app.Getenv(envName)
		// 	log.Print(envVal)
		// }
	}
}

// MovieSpinner API godoc
// note on docs: https://github.com/swaggo/swag/issues/2045#issuecomment-3892744315
//
//	@title			Bad Movie Spinner API
//	@version		1.0
//	@description	Backend API for communication with Google Sheets list of movies and TMDB
//	@host			localhost:8080
//	@BasePath		/
func main() {
	movie_api := api.NewAPIServer()

	log.Printf("Movie spinner API listening on %s", movie_api.Addr)
	if err := movie_api.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
