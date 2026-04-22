package data

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data/external"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"bytes"
	"image/jpeg"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type MovieDataHandler struct {
	spinnerAPI *external.SpinnerAPI

	// TODO: TTL caches
	imageCache     map[int]*ebiten.Image
	movieDataCache []models.MovieMeta
	logger         *slog.Logger
}

func NewMovieDataHandler(spinnerAPI *external.SpinnerAPI, logger *slog.Logger) *MovieDataHandler {
	return &MovieDataHandler{spinnerAPI: spinnerAPI, imageCache: make(map[int]*ebiten.Image), movieDataCache: make([]models.MovieMeta, 0), logger: logger}
}

func (m *MovieDataHandler) GetMovieList() []models.MovieMeta {
	movieList, err := m.spinnerAPI.GetMovies()
	if err != nil {
		return make([]models.MovieMeta, 0)
	}

	m.movieDataCache = movieList
	return movieList
}

func (m *MovieDataHandler) GetMoviePoster(tmdbID int) *ebiten.Image {
	if img, ok := m.imageCache[tmdbID]; ok {
		return img
	}

	rawPosterBytes, err := m.spinnerAPI.GetMoviePoster(tmdbID)
	if err != nil {
		return nil
	}

	rawPosterJpeg, err := jpeg.Decode(bytes.NewReader(rawPosterBytes))
	if err != nil {
		m.logger.Error("failed to decode poster image", "tmdbID", tmdbID)
		return nil
	}

	poster := ebiten.NewImageFromImage(rawPosterJpeg)
	m.imageCache[tmdbID] = poster
	return m.imageCache[tmdbID]
}
