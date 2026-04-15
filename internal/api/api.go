package api

import (
	"net/http"
	"time"
)

func NewAPIServer() *http.Server {
	router := NewRouter()

	return &http.Server{
		Handler:           router,
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
	}
}
