package main

/*
 * Simple server to host the static website containing the WASM build of spinner
 */

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/web"
	"NickBrisebois/BadMovieSpinnerGo/pkg/config"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	webConfig := &web.WebConfig{}
	if err := config.LoadConfig(webConfig, &config.ConfigOptions{Logger: logger}); err != nil {
		logger.Error("could not process env config", "err", err.Error())
		os.Exit(1)
	}

	webServer, err := web.NewWebServer(webConfig, logger)
	if err != nil {
		logger.Error("could not create web server", "err", err.Error())
		os.Exit(1)
	}
	if err := webServer.ListenAndServe(); err != nil {
		logger.Error("could not start web server", "err", err.Error())
		os.Exit(1)
	}
}
