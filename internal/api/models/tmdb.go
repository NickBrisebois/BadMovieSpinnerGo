package models

import "encoding/json"

type TMDBMovieDetails struct {
	PosterPath  string
	Overview    string
	Title       string
	Tagline     string
	VoteAverage float32
}

func (m *TMDBMovieDetails) FromJSON(data []byte) error {
	rawJSON := make(map[string]any)
	if err := json.Unmarshal(data, &rawJSON); err != nil {
		return err
	}

	m.PosterPath = rawJSON["poster_path"].(string)
	m.Overview = rawJSON["overview"].(string)
	m.Title = rawJSON["title"].(string)
	m.Tagline = rawJSON["tagline"].(string)
	m.VoteAverage = rawJSON["vote_average"].(float32)
	return nil
}
