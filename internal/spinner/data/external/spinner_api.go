package external

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const (
	getMovieURLPath  = "/sheets/movies"
	getPosterURLPath = "/sheets/movies/%d/poster"
)

type SpinnerAPI struct {
	baseURL string
	logger  *slog.Logger
}

func NewSpinnerAPI(baseURL string, logger *slog.Logger) *SpinnerAPI {
	return &SpinnerAPI{baseURL: baseURL, logger: logger}
}

func (s *SpinnerAPI) httpReq(method, path string, body io.Reader) (*http.Response, error) {
	url := s.baseURL + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func (s *SpinnerAPI) GetMovies() ([]models.MovieMeta, error) {
	s.logger.Info("fetching movie list")
	path := getMovieURLPath
	resp, err := s.httpReq("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var moviesData []models.MovieMeta
	err = json.NewDecoder(resp.Body).Decode(&moviesData)
	if err != nil {
		return nil, err
	}
	return moviesData, nil
}

func (s *SpinnerAPI) GetMoviePoster(tmdbID int) ([]byte, error) {
	path := fmt.Sprintf(getPosterURLPath, tmdbID)
	resp, err := s.httpReq("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
