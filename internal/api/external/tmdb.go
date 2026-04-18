package external

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/dto"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const (
	Lang                  = "en-US"
	TMDBGetMovieURL       = "https://api.themoviedb.org/3/movie/%s?&language=%s"
	TMDBGetMoviePosterURL = "https://image.tmdb.org/t/p/w500%s"
)

type TMDBApi struct {
	tmdbAPIKey string
	logger     *slog.Logger
}

func NewTMDBApi(tmdbAPIKey string, logger *slog.Logger) *TMDBApi {
	return &TMDBApi{
		tmdbAPIKey: tmdbAPIKey,
		logger:     logger,
	}
}

func (t *TMDBApi) httpReq(method, url string, body io.Reader) (*http.Response, error) {
	t.logger.Debug("httpReq", "method", method, "url", url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.tmdbAPIKey))
	return http.DefaultClient.Do(req)
}

func (t *TMDBApi) FetchMovieData(tmdbID string) (*dto.TMDBMovieDetailResponse, error) {
	getURL := fmt.Sprintf(TMDBGetMovieURL, tmdbID, Lang)
	resp, err := t.httpReq("GET", getURL, nil)
	if err != nil {
		t.logger.Error("failed to fetch movie data", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	var movieDetails dto.TMDBMovieDetailResponse
	err = json.NewDecoder(resp.Body).Decode(&movieDetails)
	if err != nil {
		return nil, err
	}

	return &movieDetails, nil
}

func (t *TMDBApi) FetchMoviePoster(posterPath string) ([]byte, error) {
	getURL := fmt.Sprintf(TMDBGetMoviePosterURL, posterPath)
	resp, err := t.httpReq("GET", getURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
