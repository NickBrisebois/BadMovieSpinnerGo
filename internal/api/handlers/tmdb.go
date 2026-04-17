package handlers

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/dto"
	"NickBrisebois/BadMovieSpinnerGo/internal/api/external"
	"log/slog"
	"net/url"
	"path"
)

type TMDBHandler struct {
	tmdbAPI *external.TMDBApi
	logger  *slog.Logger

	tmdbMovieCache map[string]dto.TMDBMovieDetails
}

func NewTMDBHandler(tmdbAccessToken string, logger *slog.Logger) *TMDBHandler {
	tmdbAPI := external.NewTMDBApi(tmdbAccessToken, logger)
	return &TMDBHandler{tmdbAPI: tmdbAPI, logger: logger}
}

func (h *TMDBHandler) GetMovieIDFromURL(movieURL string) (string, error) {
	parsedURL, err := url.Parse(movieURL)
	if err != nil {
		return "", err
	}

	return path.Base(parsedURL.Path), nil
}

func (h *TMDBHandler) GetMovieData(tmdbID string) (*dto.TMDBMovieDetails, error) {
	if movieData, ok := h.tmdbMovieCache[tmdbID]; ok {
		return &movieData, nil
	}

	fetchedMovieData, err := h.tmdbAPI.FetchMovieData(tmdbID)
	if err != nil {
		return nil, err
	}

	h.tmdbMovieCache[tmdbID] = *fetchedMovieData
	return fetchedMovieData, nil
}
