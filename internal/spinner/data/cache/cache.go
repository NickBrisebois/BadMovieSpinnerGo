package cache

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
)

type movieListCachePayload struct {
	// movieListCachePayload represents the structure written to a cache file
	Movies []models.MovieMeta
}

type Cache interface {
	GetMoviePoster(key int) ([]byte, bool)
	PutMoviePoster(key int, data []byte) error
	GetMovieList() ([]models.MovieMeta, error)
	PutMovieList(movies []models.MovieMeta) error
}
