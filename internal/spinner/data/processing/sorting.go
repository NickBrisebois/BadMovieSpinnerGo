package processing

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"errors"
)

type SortType int

const (
	SortName SortType = iota
	SortSuggestedBy
)

type MovieSortOptions struct {
	Type SortType
}

func SortMovieList(movies []models.MovieMeta, options *MovieSortOptions) (map[string][]models.MovieMeta, error) {
	sorted := make(map[string][]models.MovieMeta)
	if options == nil {
		options = &MovieSortOptions{Type: SortName}
	}

	switch options.Type {
	case SortName:
		for _, movie := range movies {
			sorted[movie.Title] = append(sorted[movie.Title], movie)
		}
	case SortSuggestedBy:
		for _, movie := range movies {
			sorted[movie.SuggestedBy] = append(sorted[movie.SuggestedBy], movie)
		}
	default:
		return nil, errors.New("invalid sort type")
	}
	return sorted, nil
}
