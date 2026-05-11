package web

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
)

//go:embed html
var spinnerRootFS embed.FS

func NewWebServer(config *WebConfig, logger *slog.Logger) (*http.Server, error) {
	mux := http.NewServeMux()

	// just serve the exact embedded html static assets folder on the root
	webDir, err := fs.Sub(spinnerRootFS, "html")
	if err != nil {
		logger.Error("failed to get web root fs", "err", err)
		return nil, err
	}
	webRoot := http.FileServer(http.FS(webDir))
	mux.Handle("/", webRoot)

	serverAddr := config.WebHost + ":" + config.WebPort
	logger.Debug("web server listening on", "addr", serverAddr)
	return &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}, nil
}
