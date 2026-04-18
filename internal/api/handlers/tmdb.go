package handlers

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/dto"
	"NickBrisebois/BadMovieSpinnerGo/internal/api/external"
	"log/slog"
	"net/url"
	"path"
	"sync"
)

type TMDBHandler struct {
	tmdbAPI      *external.TMDBApi
	imageHandler *ImageHandler
	logger       *slog.Logger

	tmdbMovieCache map[string]dto.TMDBMovieDetailResponse
}

func NewTMDBHandler(tmdbAccessToken string, imageHandler *ImageHandler, logger *slog.Logger) *TMDBHandler {
	tmdbAPI := external.NewTMDBApi(tmdbAccessToken, logger)
	return &TMDBHandler{
		tmdbAPI:        tmdbAPI,
		imageHandler:   imageHandler,
		logger:         logger,
		tmdbMovieCache: make(map[string]dto.TMDBMovieDetailResponse),
	}
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

func (t *TMDBHandler) BulkFetchMovieData(tmdbIDs []string) (*[]dto.TMDBMovieDetailResponse, error) {
	movieDetailsChan := make(chan *dto.TMDBMovieDetailResponse)

	// tmdb doesn't seem to have a bulk movie API so we have to hit them a bunch of times :(
	// to speed things up for when we don't have the data cached, we request each movie concurrently
	var reqWaitGroup sync.WaitGroup
	for _, id := range tmdbIDs {
		reqWaitGroup.Go(func() {
			movieData, err := t.GetMovieData(id)
			if err != nil {
				t.logger.Error("failed to fetch movie data", "tmdbID", id, "error", err)
				return
			}
			movieDetailsChan <- movieData
		})
	}
	reqWaitGroup.Wait()

	close(movieDetailsChan)
	var movieDetails []dto.TMDBMovieDetailResponse
	for movieData := range movieDetailsChan {
		movieDetails = append(movieDetails, *movieData)
	}

	return &movieDetails, nil
}

func (h *TMDBHandler) GetMoviePoster(tmdbID string) ([]byte, error) {
	// Check cache for the image before we make any requests
	posterImgData, found := h.imageHandler.GetImage(tmdbID)
	if found {
		return posterImgData, nil
	}

	// this isn't ideal, but, lookup the movie data to get the movie poster path.
	// if this method is being called then it's more likely than not that we've
	// already cached the data so this isn't expensive
	movieData, err := h.GetMovieData(tmdbID)
	if err != nil {
		return nil, err
	}

	// If no cached image then we must resort to fetching it from TMDB :(
	fetchedPosterImg, err := h.tmdbAPI.FetchMoviePoster(movieData.PosterPath)
	if err != nil {
		return nil, err
	}

	h.imageHandler.CacheImage(fetchedPosterImg, tmdbID)
	return fetchedPosterImg, nil
}
