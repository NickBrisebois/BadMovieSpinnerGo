package models

type MovieMeta struct {
	Title       string
	TMDBId      int
	Watched     bool
	SuggestedBy string
	PosterURL   string
	Description *string
}
