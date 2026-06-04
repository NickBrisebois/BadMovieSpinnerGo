package models

type PersonName string

type MovieMeta struct {
	Title       string
	TMDBId      int
	Watched     bool
	SuggestedBy PersonName
	PosterURL   string
	Description *string
}
