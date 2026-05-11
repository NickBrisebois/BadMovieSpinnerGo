package api

type Config struct {
	Server struct {
		Host           string   `env:"SERVER_HOST" default:"localhost"`
		Port           string   `env:"SERVER_PORT" default:"8080"`
		AllowedOrigins []string `env:"ALLOWED_ORIGINS" delimiter:","`
	}
	Auth struct {
		GCPServiceAccountKeyPath string `env:"GOOGLE_SERVICE_ACCOUNT_KEY_PATH"`
		GCPScopes                string `env:"GOOGLE_SCOPES"`
		TMDBReadAccessToken      string `env:"TMDB_ACCESS_TOKEN"`
	}
	GSheets struct {
		SheetID string `env:"GOOGLE_SHEET_ID"`
	}
	Cache struct {
		ImageCacheDir string `env:"IMAGE_CACHE_DIR"`
	}
}
