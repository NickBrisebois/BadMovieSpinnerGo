package data

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/cache"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/external"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/processing"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"bytes"
	"image/jpeg"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type GetMovieListOptions struct {
	Filters *processing.MovieFilters
}

type MovieDataHandler struct {
	spinnerAPI *external.SpinnerAPI

	memImageCache     map[int]*ebiten.Image
	memMovieDataCache []models.MovieMeta

	cache cache.Cache

	logger *slog.Logger
}

func NewMovieDataHandler(spinnerAPI *external.SpinnerAPI, logger *slog.Logger) *MovieDataHandler {
	cache, err := cache.NewCache(logger)
	if err != nil {
		logger.Error("failed to create cache", "error", err)
		return nil
	}
	return &MovieDataHandler{
		spinnerAPI:        spinnerAPI,
		memImageCache:     make(map[int]*ebiten.Image),
		memMovieDataCache: make([]models.MovieMeta, 0),
		cache:             cache,
		logger:            logger,
	}
}

func (m *MovieDataHandler) GetMovieList(options *GetMovieListOptions) []models.MovieMeta {

	// check if we have the movie list cached in mem
	if len(m.memMovieDataCache) == 0 {
		// if not, check if we have it in the FS or localstorage
		movieList, err := m.cache.GetMovieList()
		if err != nil {
			// if neither caches have what we need, resort to pulling it from the API
			movieList, err = m.spinnerAPI.GetMovies()
			if err != nil {
				m.logger.Error("failed to get movie list from API", "error", err)
				return make([]models.MovieMeta, 0)
			}
			m.cache.PutMovieList(movieList)
		}
		m.memMovieDataCache = movieList
	}

	if options != nil && options.Filters != nil {
		return processing.FilterMovieList(m.memMovieDataCache, options.Filters)
	}

	return m.memMovieDataCache
}

func (m *MovieDataHandler) GetMoviesBySuggestedBy(options *GetMovieListOptions) map[string][]models.MovieMeta {
	movieList := m.GetMovieList(options)
	sortedMovies, err := processing.SortMovieList(movieList, &processing.MovieSortOptions{Type: processing.SortSuggestedBy})
	if err != nil {
		m.logger.Error("failed to sort movie list", "error", err)
		return nil
	}
	return sortedMovies
}

func (m *MovieDataHandler) GetMoviePoster(tmdbID int) *ebiten.Image {

	// first check if we have the poster in mem
	if img, ok := m.memImageCache[tmdbID]; ok {
		return img
	}

	// second, check if we have the poster cached in the filesystem/browser
	rawPosterBytes, ok := m.cache.GetMoviePoster(tmdbID)
	if !ok {
		var err error
		rawPosterBytes, err = m.spinnerAPI.GetMoviePoster(tmdbID)
		if err != nil {
			return nil
		}

		err = m.cache.PutMoviePoster(tmdbID, rawPosterBytes)
		if err != nil {
			m.logger.Error("failed to cache poster image", "tmdbID", tmdbID, "error", err)
		}
	}

	rawPosterJpeg, err := jpeg.Decode(bytes.NewReader(rawPosterBytes))
	if err != nil {
		m.logger.Error("failed to decode poster image", "tmdbID", tmdbID)
		return nil
	}

	poster := ebiten.NewImageFromImage(rawPosterJpeg)
	m.memImageCache[tmdbID] = poster
	return m.memImageCache[tmdbID]
}
