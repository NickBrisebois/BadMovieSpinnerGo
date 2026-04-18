package api

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Hacky Config manager that loads all configuration from env variables
type Config struct {
	Server struct {
		Host string `env:"SERVER_HOST"`
		Port int    `env:"SERVER_PORT"`
	}
	Auth struct {
		GCPServiceAccountKeyPath string `env:"GOOGLE_SERVICE_ACCOUNT_KEY_PATH"`
		GCPScopes                string `env:"GOOGLE_SCOPES"`
		TMDBApiKey               string `env:"TMDB_API_KEY"`
		TMDBReadAccessToken      string `env:"TMDB_ACCESS_TOKEN"`
	}
	GSheets struct {
		SheetID string `env:"GOOGLE_SHEET_ID"`
	}
	Cache struct {
		ImageCacheDir string `env:"IMAGE_CACHE_DIR"`
	}
}

func LoadConfig(conf *Config) error {
	valOfConfig := reflect.ValueOf(conf)
	return parseEnvsIntoConfig(valOfConfig.Elem())
}

// parseEnvsIntoConfig recursively digs into given struct type and fills in properties with env values
// keyed by the tag provided in the struct declaration
func parseEnvsIntoConfig(conf reflect.Value) error {
	confType := conf.Type()

	for i := 0; i < conf.NumField(); i++ {
		confProperty := conf.Field(i)
		confPropertyTags := confType.Field(i)

		if confProperty.Kind() == reflect.Struct {
			if err := parseEnvsIntoConfig(confProperty); err != nil {
				return err
			}

			continue
		}

		tag := confPropertyTags.Tag.Get("env")
		parts := strings.Split(tag, ",")
		key := strings.TrimSpace(parts[0])

		raw, ok := os.LookupEnv(key)
		if !ok || raw == "" {
			return fmt.Errorf("missing env var %q", key)
		}

		switch confProperty.Kind() {
		case reflect.String:
			confProperty.SetString(raw)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(raw, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int value for %q: %w", key, err)
			}
			confProperty.SetInt(n)
		case reflect.Bool:
			b, err := strconv.ParseBool(raw)
			if err != nil {
				return fmt.Errorf("invalid bool value for %q: %w", key, err)
			}
			confProperty.SetBool(b)
		default:
			return fmt.Errorf("unsupported kind %v for %q", conf.Kind(), key)
		}
	}

	return nil
}
