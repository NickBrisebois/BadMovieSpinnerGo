package models

type Movie struct {
	Title       string
	Link        string
	Watched     bool
	SuggestedBy string
	PosterURL   string
	Description *string
}
