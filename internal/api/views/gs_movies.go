package views

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/handlers"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type GSheetsView struct {
	gsheetsHandler *handlers.GSMoviesHandler
	imageHandler   *handlers.ImageHandler
	tmdbHandler    *handlers.TMDBHandler
	logger         *slog.Logger
}

func NewGSheetsView(credentialsFilePath, spreadsheetID, tmdbAccessToken, imageCacheDir string, logger *slog.Logger) (*GSheetsView, error) {
	imageHandler := handlers.NewImageHandler(imageCacheDir, logger)
	gsheetsHandler, err := handlers.NewGSheetsHandler(credentialsFilePath, spreadsheetID, tmdbAccessToken, imageHandler, logger)
	tmdbHandler := handlers.NewTMDBHandler(tmdbAccessToken, imageHandler, logger)
	if err != nil {
		return nil, err
	}
	return &GSheetsView{
		gsheetsHandler: gsheetsHandler,
		imageHandler:   imageHandler,
		tmdbHandler:    tmdbHandler,
		logger:         logger,
	}, nil
}

// GetMovieList
//
//	@Summary	Handle HTTP request for retrieving list of movies from Google Sheets (specified by local spreadsheetID).
//	@Tags		movies
//	@Produce	json
//	@Success	200	{array}	models.MovieMeta
//	@Router		/sheets/movies [get]
func (h *GSheetsView) GetMovieList(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("GetMovieList requested")
	movies, err := h.gsheetsHandler.GetAllMovies()
	if err != nil {
		h.logger.Error("failed to retrieve movie data", "err", err)
	}

	resp_json, err := json.Marshal(movies)
	if err != nil {
		h.logger.Error("failed to marshal movie data", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp_json)
}

// GetMoviePoster
//
//	@Summary	Handle HTTP request for retrieving a movie poster from TMDB (specified by TMDB ID).
//	@Tags		movies
//	@Produce	image/jpeg
//	@Success	200	{file}	file	"movie poster"
//	@Router		/sheets/movies/{tmdbID}/poster [get]
func (h *GSheetsView) GetMoviePoster(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("GetMoviePoster requested")
	strTMDBId := r.PathValue("tmdbID")
	tmdbID, err := strconv.Atoi(strTMDBId)
	if strTMDBId == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid TMDB movie ID"))
		return
	}

	posterImgData, err := h.tmdbHandler.GetMoviePoster(tmdbID)
	if err != nil {
		h.logger.Error("failed to retrieve movie poster", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(posterImgData)
}
