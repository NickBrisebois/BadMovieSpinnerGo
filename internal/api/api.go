package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func NewAPIServer(config *Config, logger *slog.Logger) *http.Server {
	router := NewRouter(config, logger)
	serverAddr := fmt.Sprintf("%s:%s", config.Server.Host, strconv.Itoa(config.Server.Port))
	return &http.Server{
		Handler:           router,
		Addr:              serverAddr,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
