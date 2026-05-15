package processing

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
)

type WatchedStatus int

const (
	WatchedStatusAny WatchedStatus = iota
	WatchedStatusWatched
	WatchedStatusUnwatched
)

type MovieFilters struct {
	Watched WatchedStatus
}

func FilterWatched(movieList []models.MovieMeta) []models.MovieMeta {
	var filtered []models.MovieMeta
	for _, movie := range movieList {
		if movie.Watched {
			filtered = append(filtered, movie)
		}
	}
	return filtered
}

func FilterUnwatched(movieList []models.MovieMeta) []models.MovieMeta {
	var filtered []models.MovieMeta
	for _, movie := range movieList {
		if !movie.Watched {
			filtered = append(filtered, movie)
		}
	}
	return filtered
}

func FilterMovieList(movieList []models.MovieMeta, movieFilters *MovieFilters) []models.MovieMeta {
	if movieFilters == nil {
		return movieList
	}

	filtered := movieList
	switch movieFilters.Watched {
	case WatchedStatusWatched:
		filtered = FilterWatched(filtered)
	case WatchedStatusUnwatched:
		filtered = FilterUnwatched(filtered)
	}

	return filtered
}
