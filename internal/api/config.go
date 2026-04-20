package api

type Config struct {
	Server struct {
		Host string `env:"SERVER_HOST" default:"localhost"`
		Port int    `env:"SERVER_PORT" default:"8080"`
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
