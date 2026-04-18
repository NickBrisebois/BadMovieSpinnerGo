package handlers

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/dto"
	"NickBrisebois/BadMovieSpinnerGo/internal/api/external"
	"log/slog"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"
)

type TMDBHandler struct {
	tmdbAPI      *external.TMDBApi
	imageHandler *ImageHandler
	logger       *slog.Logger

	tmdbMovieCache map[int]dto.TMDBMovieDetailResponse
	cacheMu        sync.RWMutex
}

func NewTMDBHandler(tmdbAccessToken string, imageHandler *ImageHandler, logger *slog.Logger) *TMDBHandler {
	tmdbAPI := external.NewTMDBApi(tmdbAccessToken, logger)
	return &TMDBHandler{
		tmdbAPI:        tmdbAPI,
		imageHandler:   imageHandler,
		logger:         logger,
		tmdbMovieCache: make(map[int]dto.TMDBMovieDetailResponse),
	}
}

func (h *TMDBHandler) getCacheItem(tmdbID int) (*dto.TMDBMovieDetailResponse, bool) {
	h.cacheMu.RLock()
	if movieData, ok := h.tmdbMovieCache[tmdbID]; ok {
		defer h.cacheMu.RUnlock()
		return &movieData, true
	}
	defer h.cacheMu.RUnlock()
	return nil, false
}

func (h *TMDBHandler) setCacheItem(tmdbID int, movieDetailResponse *dto.TMDBMovieDetailResponse) {
	h.cacheMu.Lock()
	h.tmdbMovieCache[tmdbID] = *movieDetailResponse
	h.cacheMu.Unlock()
}

func (h *TMDBHandler) GetMovieIDFromURL(movieURL string) (int, error) {
	parsedURL, err := url.Parse(movieURL)
	if err != nil {
		return -1, err
	}

	// get the named TMDB ID from the URL (ie. "950387-a-minecraft-movie") then pull the int ID out
	namedTMDBID := path.Base(parsedURL.Path)
	strSplitIndex := strings.Index(namedTMDBID, "-")
	if strSplitIndex == -1 {
		// possibly already just an int so convert and return
		parsedID, err := strconv.Atoi(namedTMDBID)
		return parsedID, err
	}

	tmdbID := namedTMDBID[:strSplitIndex]
	parsedID, err := strconv.Atoi(tmdbID)
	if err != nil {
		return -1, err
	}
	return parsedID, err
}

func (h *TMDBHandler) GetMovieData(tmdbID int) (*dto.TMDBMovieDetailResponse, error) {
	if movieData, ok := h.getCacheItem(tmdbID); ok {
		return movieData, nil
	}

	fetchedMovieData, err := h.tmdbAPI.FetchMovieData(tmdbID)
	if err != nil {
		return nil, err
	}

	h.setCacheItem(tmdbID, fetchedMovieData)
	return fetchedMovieData, nil
}

func (t *TMDBHandler) BulkFetchMovieData(tmdbIDs []int) (map[int]dto.TMDBMovieDetailResponse, error) {
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

	go func() {
		reqWaitGroup.Wait()
		close(movieDetailsChan)
	}()

	movieDetails := make(map[int]dto.TMDBMovieDetailResponse)
	for movieData := range movieDetailsChan {
		movieDetails[movieData.ID] = *movieData
	}

	return movieDetails, nil
}

func (h *TMDBHandler) GetMoviePoster(tmdbID int) ([]byte, error) {
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
