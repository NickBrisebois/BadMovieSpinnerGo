package api

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/handlers"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handlers.GetHealthz)

	return mux
}
