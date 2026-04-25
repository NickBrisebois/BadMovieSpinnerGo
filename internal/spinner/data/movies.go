package data

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/cache"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/external"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"bytes"
	"image/jpeg"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func (m *MovieDataHandler) GetMovieList() []models.MovieMeta {
	var movieList []models.MovieMeta
	var err error

	if movieList, err = m.cache.GetMovieList(); err != nil {
		movieList, err = m.spinnerAPI.GetMovies()
		if err != nil {
			return make([]models.MovieMeta, 0)
		}
		m.cache.PutMovieList(movieList)
	}

	m.memMovieDataCache = movieList
	return movieList
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
