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
	handler := &GSheetsMovieReqHandler{sheetsAPI: nil}
	sheetsAPI, err := external.NewGoogleSheetsAPI(
		credentialsFilePath,
		spreadsheetID,
		logger,
	)
	if err != nil {
		return nil, err
	}
	handler.sheetsAPI = sheetsAPI
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
		h.logger.Error("Failed to retrieve movie data", "err", err)
	}

	h.logger.Debug("movies", "movies", movies)
}
