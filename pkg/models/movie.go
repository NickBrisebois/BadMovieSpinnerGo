package models

type MovieMeta struct {
	Title       string
	TMDBId      string
	Watched     bool
	SuggestedBy string
	PosterURL   string
	Description *string
}
