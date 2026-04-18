package external

import (
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"bytes"
	"io"
	"net/http"
)

const (
	getMovieURLPath  = "/movies"
	getPosterURLPath = "/movies/%s/poster"
)

type SpinnerAPI struct {
	baseURL string
}

func NewSpinnerAPI(baseURL string) *SpinnerAPI {
	return &SpinnerAPI{baseURL: baseURL}
}

func (s *SpinnerAPI) httpReq(method, path string, body []byte) ([]byte, error) {
	url := s.baseURL + path
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (s *SpinnerAPI) GetMovies() ([]models.MovieMeta, error) {
	return "", nil
}

func (s *SpinnerAPI) GetMoviePoster(tmdbID string) ([]byte, error) {

}
