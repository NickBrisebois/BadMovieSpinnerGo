package external

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetsAPI struct {
	sheetsService *sheets.Service
	spreadsheetID string
	rawGoogleDoc  *sheets.Spreadsheet
	logger        *slog.Logger
}

func NewGoogleSheetsAPI(credentialsFilePath, spreadsheetID string, logger *slog.Logger) (*GoogleSheetsAPI, error) {
	sheetsClientOptions := option.WithAuthCredentialsFile(option.ServiceAccount, credentialsFilePath)
	sheetsService, err := sheets.NewService(
		context.Background(),
		sheetsClientOptions,
		option.WithScopes(sheets.SpreadsheetsReadonlyScope),
	)
	if err != nil {
		return nil, err
	}

	spreadsheet, err := sheetsService.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, err
	}
	return &GoogleSheetsAPI{sheetsService: sheetsService, spreadsheetID: spreadsheetID, rawGoogleDoc: spreadsheet, logger: logger}, nil
}

func (g *GoogleSheetsAPI) GetMovieData() (*[]string, error) {
	if len(g.rawGoogleDoc.Sheets) == 0 {
		g.logger.Error("No Google Sheet found for spreadsheet ID", "spreadSheetID", g.spreadsheetID)
		return nil, nil
	}

	mainSheetName := g.rawGoogleDoc.Sheets[0].Properties.Title
	readRange := fmt.Sprintf("%s!A:Z", mainSheetName)

	movieListValues, err := g.sheetsService.Spreadsheets.Values.Get(g.spreadsheetID, readRange).Do()
	if err != nil {
		return nil, nil
	}

	movies := make([]string, 0, len(movieListValues.Values))
	for _, row := range movieListValues.Values {
		if len(row) > 0 {
			movies = append(movies, row[0].(string))
		}
	}
	return &movies, nil
}
