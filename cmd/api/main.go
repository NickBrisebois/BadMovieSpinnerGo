package main

import (
	"log"
	"reflect"
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

	log.Println("I'm a teapot")
}
