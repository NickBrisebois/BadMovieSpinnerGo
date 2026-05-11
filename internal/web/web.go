package web

import (
	"embed"
	"log/slog"
	"net/http"
)

//go:embed html
var spinnerRootFS embed.FS

func NewWebServer(config *WebConfig, logger *slog.Logger) (*http.Server, error) {
	mux := http.NewServeMux()

	// just serve the exact embedded html static assets folder on the root
	mux.Handle("/", http.FileServer(http.FS(spinnerRootFS)))

	serverAddr := config.WebHost + ":" + config.WebPort
	logger.Debug("web server listening on", "addr", serverAddr)
	return &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}, nil
}
