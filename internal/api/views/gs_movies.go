package views

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/handlers"
	"log/slog"
	"net/http"
)

type GSheetsView struct {
	gsheetsHandler *handlers.GSheetsHandler
	tmdbHandler    *handlers.TMDBHandler
	logger         *slog.Logger
}

func NewGSheetsView(credentialsFilePath, spreadsheetID, tmdbAccessToken string, logger *slog.Logger) (*GSheetsView, error) {
	gsheetsHandler, err := handlers.NewGSheetsHandler(credentialsFilePath, spreadsheetID, tmdbAccessToken, logger)
	if err != nil {
		return nil, err
	}
	return &GSheetsView{
		gsheetsHandler: gsheetsHandler,
		logger:         logger,
	}, nil
}

// GetMovieList
//
//	@Summary	Handle HTTP request for retrieving list of movies from Google Sheets (specified by local spreadsheetID).
//	@Tags		movies
//	@Produce	json
//	@Success	200	{array}	string
//	@Router		/sheets/movies [get]
func (h *GSheetsView) GetMovieList(w http.ResponseWriter, r *http.Request) {
	movies, err := h.gsheetsHandler.GetAllMovies()
	if err != nil {
		h.logger.Error("failed to retrieve movie data", "err", err)
	}

	h.logger.Debug("movies", "movies", movies)
}
