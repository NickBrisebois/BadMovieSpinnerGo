package data

import "NickBrisebois/BadMovieSpinnerGo/pkg/models"

func PullMovieList() []models.MovieMeta {
	// TODO: Replace with an actual API call
	return []models.MovieMeta{
		{Title: "Movie 1", Watched: false, SuggestedBy: "user1", PosterURL: "https://http.cat/images/100.jpg", TMDBId: "316776-who-killed-captain-alex"},
		{Title: "Movie 2", Watched: true, SuggestedBy: "user2", PosterURL: "https://http.cat/images/101.jpg", TMDBId: "316776-who-killed-captain-alex"},
		{Title: "Movie 3", Watched: false, SuggestedBy: "user3", PosterURL: "https://http.cat/images/102.jpg", TMDBId: "316776-who-killed-captain-alex"},
	}
}
