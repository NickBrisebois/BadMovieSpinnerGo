package handlers

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/dto"
	"NickBrisebois/BadMovieSpinnerGo/internal/api/external"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"log/slog"
)

type GSMoviesHandler struct {
	sheetsAPI   *external.GoogleSheetsAPI
	tmdbHandler *TMDBHandler
	logger      *slog.Logger
}

func NewGSheetsHandler(
	credentialsFilePath, spreadsheetID, tmdbAccessToken string,
	imageHandler *ImageHandler,
	logger *slog.Logger,

) (*GSMoviesHandler, error) {

	sheetsAPI, err := external.NewGoogleSheetsAPI(
		credentialsFilePath,
		spreadsheetID,
		logger,
	)
	if err != nil {
		logger.Error("failed to create google sheets api", "err", err)
		return nil, err
	}

	tmdbHandler := NewTMDBHandler(tmdbAccessToken, imageHandler, logger)
	handler := &GSMoviesHandler{
		sheetsAPI:   sheetsAPI,
		tmdbHandler: tmdbHandler,
		logger:      logger,
	}
	return handler, nil
}

func (h *GSMoviesHandler) enrichMovieList(rawMovies []dto.GSheetsMoviesEntry) ([]models.MovieMeta, error) {
	var movies []models.MovieMeta

	// pull all the TMDB IDs out of the list based on the TMDB link data
	tmdbIDs := make([]int, len(rawMovies))
	for i, rawMovie := range rawMovies {
		tmdbID, err := h.tmdbHandler.GetMovieIDFromURL(rawMovie.TMDBLink)
		if err != nil {
			h.logger.Error("failed to get movie ID from URL", "tmdbLink", rawMovie.TMDBLink, "error", err)
		}
		tmdbIDs[i] = tmdbID
	}

	// bulk fetch all the movies (this is a request per movie but it's batched with goroutines)
	bulkMovies, err := h.tmdbHandler.BulkFetchMovieData(tmdbIDs)
	if err != nil {
		h.logger.Error("failed to bulk fetch movie data", "error", err)
		return nil, err
	}

	for _, enrichedMovie := range bulkMovies {
		movies = append(movies, models.MovieMeta{
			Title:       enrichedMovie.Title,
			TMDBId:      enrichedMovie.ID,
			Watched:     false,
			SuggestedBy: "nick",
			PosterURL:   enrichedMovie.PosterPath,
			Description: &enrichedMovie.Overview,
		})
	}
	return movies, nil
}

func (h *GSMoviesHandler) GetAllMovies() ([]models.MovieMeta, error) {
	movies, err := h.sheetsAPI.GetMovieData()
	if err != nil {
		return nil, err
	}

	enrichedMoviesList, err := h.enrichMovieList(movies)
	if err != nil {
		return nil, err
	}

	return enrichedMoviesList, nil
}
