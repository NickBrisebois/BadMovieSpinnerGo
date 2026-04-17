package handlers

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/dto"
	"NickBrisebois/BadMovieSpinnerGo/internal/api/external"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"log/slog"
)

type GSheetsHandler struct {
	sheetsAPI   *external.GoogleSheetsAPI
	tmdbHandler *TMDBHandler
	logger      *slog.Logger
}

func NewGSheetsHandler(
	credentialsFilePath, spreadsheetID, tmdbAccessToken string,
	logger *slog.Logger,

) (*GSheetsHandler, error) {

	sheetsAPI, err := external.NewGoogleSheetsAPI(
		credentialsFilePath,
		spreadsheetID,
		logger,
	)
	if err != nil {
		logger.Error("failed to create google sheets api", "err", err)
		return nil, err
	}

	tmdbHandler := NewTMDBHandler(tmdbAccessToken, logger)
	handler := &GSheetsHandler{
		sheetsAPI:   sheetsAPI,
		tmdbHandler: tmdbHandler,
		logger:      logger,
	}
	return handler, nil
}

func (h *GSheetsHandler) enrichMovieList(rawMovies []dto.GSheetsMoviesEntry) ([]models.MovieMeta, error) {
	var movies []models.MovieMeta
	for _, movie := range rawMovies {
		tmdbID, err := h.tmdbHandler.GetMovieIDFromURL(movie.TMDBLink)
		if err != nil {
			h.logger.Error("failed to get movie ID from URL", "error", err)
			continue
		}

		tmdbData, err := h.tmdbHandler.GetMovieData(*tmdbID)
		if err != nil {
			return nil, err
		}
		movies = append(movies, models.MovieMeta{
			Title:       tmdbData.Title,
			TMDBId:      *tmdbID,
			Watched:     movie.Watched,
			SuggestedBy: movie.SuggestedBy,
			PosterURL:   tmdbData.PosterPath,
			Description: &tmdbData.Overview,
		})
	}
	return movies, nil
}

func (h *GSheetsHandler) GetAllMovies() ([]models.MovieMeta, error) {
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
