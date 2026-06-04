package processing

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
)

func SortMovieListByPerson(movies []models.MovieMeta) map[models.PersonName][]models.MovieMeta {
	sorted := make(map[models.PersonName][]models.MovieMeta)
	for _, movie := range movies {
		sorted[movie.SuggestedBy] = append(sorted[movie.SuggestedBy], movie)
	}
	return sorted
}
