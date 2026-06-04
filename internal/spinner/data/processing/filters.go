package processing

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"slices"
)

type WatchedStatus int

const (
	WatchedStatusAny WatchedStatus = iota
	WatchedStatusWatched
	WatchedStatusUnwatched
)

type MovieFilters struct {
	Watched     WatchedStatus
	SuggestedBy *[]models.PersonName
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

func FilterSuggestedBy(movieList []models.MovieMeta, suggestedBy []models.PersonName) []models.MovieMeta {
	var filtered []models.MovieMeta
	for _, movie := range movieList {
		if slices.Contains(suggestedBy, movie.SuggestedBy) {
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

	if movieFilters.SuggestedBy != nil {
		filtered = FilterSuggestedBy(filtered, *movieFilters.SuggestedBy)
	}

	return filtered
}
