package models

type MovieMeta struct {
	Title       string
	Link        string
	Watched     bool
	SuggestedBy string
	PosterURL   string
	Description *string
}
