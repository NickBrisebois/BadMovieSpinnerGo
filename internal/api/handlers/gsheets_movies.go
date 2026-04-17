package handlers

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/external"
	"log/slog"
	"net/http"
)

type GSheetsMovieReqHandler struct {
	sheetsAPI *external.GoogleSheetsAPI
	logger    *slog.Logger
}

func NewGSheetsReqHandler(credentialsFilePath, spreadsheetID string, logger *slog.Logger) (*GSheetsMovieReqHandler, error) {

	sheetsAPI, err := external.NewGoogleSheetsAPI(
		credentialsFilePath,
		spreadsheetID,
		logger,
	)
	if err != nil {
		logger.Error("failed to create google sheets api", "err", err)
		return nil, err
	}
	handler := &GSheetsMovieReqHandler{
		sheetsAPI: sheetsAPI,
		logger:    logger,
	}
	return handler, nil
}

// GetMovieList
//
//	@Summary	Handle HTTP request for retrieving list of movies from Google Sheets (specified by local spreadsheetID).
//	@Tags		movies
//	@Produce	json
//	@Success	200	{array}	string
//	@Router		/sheets/movies [get]
func (h *GSheetsMovieReqHandler) GetMovieList(w http.ResponseWriter, r *http.Request) {
	movies, err := h.sheetsAPI.GetMovieData()
	if err != nil {
		h.logger.Error("failed to retrieve movie data", "err", err)
	}

	h.logger.Debug("movies", "movies", movies)
}
