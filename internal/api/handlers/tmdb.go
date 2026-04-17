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

	tmdbMovieCache map[string]dto.TMDBMovieDetailResponse
}

func NewTMDBHandler(tmdbAccessToken string, logger *slog.Logger) *TMDBHandler {
	tmdbAPI := external.NewTMDBApi(tmdbAccessToken, logger)
	return &TMDBHandler{tmdbAPI: tmdbAPI, logger: logger, tmdbMovieCache: make(map[string]dto.TMDBMovieDetailResponse)}
}

func (h *TMDBHandler) GetMovieIDFromURL(movieURL string) (*string, error) {
	parsedURL, err := url.Parse(movieURL)
	if err != nil {
		return nil, err
	}

	tmdbID := path.Base(parsedURL.Path)
	return &tmdbID, nil
}

func (h *TMDBHandler) GetMovieData(tmdbID string) (*dto.TMDBMovieDetailResponse, error) {
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
