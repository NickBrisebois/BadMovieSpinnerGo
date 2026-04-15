package handlers

import "NickBrisebois/BadMovieSpinnerGo/internal/api/external"

type GSheetsHandler struct {
	sheetsAPI *external.GoogleSheetsAPI
}

func NewGSheetsHandler(credentialsFilePath, spreadsheetID string) (*GSheetsHandler, error) {
	handler := &GSheetsHandler{sheetsAPI: nil}
	sheetsAPI, err := external.NewGoogleSheetsAPI(
		credentialsFilePath,
		spreadsheetID,
	)
	if err != nil {
		return nil, err
	}
	handler.sheetsAPI = sheetsAPI
	return handler, nil
}

func (h *GSheetsHandler) GetMovieList() {

}
