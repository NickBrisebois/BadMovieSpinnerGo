package main

import (
	"log"
	"reflect"

	"NickBrisebois/BadMovieSpinnerGo/internal/api"
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

func main() {
	movie_api := api.NewAPIServer()

	log.Printf("Movie spinner API listening on %s", movie_api.Addr)
	if err := movie_api.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
