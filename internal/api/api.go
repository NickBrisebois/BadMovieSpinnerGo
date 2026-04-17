package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func NewAPIServer(config *Config, logger *slog.Logger) (*http.Server, error) {
	router, err := NewRouter(config, logger)
	if err != nil {
		logger.Error("failed to create router", slog.Any("error", err))
		return nil, err
	}
	serverAddr := fmt.Sprintf("%s:%s", config.Server.Host, strconv.Itoa(config.Server.Port))
	return &http.Server{
		Handler:           router,
		Addr:              serverAddr,
		ReadHeaderTimeout: 5 * time.Second,
	}, nil
}
