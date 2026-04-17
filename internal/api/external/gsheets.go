package external

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/api/dto"
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// The row index for where the table header is
// Used to jump to when processing data
const moviesHeaderStartIndex = 5

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

// extractMoviesFromSheetData extracts movies from the sheet data and returns array of unprocessed raw sheet data structs
func (g *GoogleSheetsAPI) extractMoviesFromSheetData(sheetData [][]any) ([]dto.GSheetsMoviesEntry, error) {
	if len(sheetData) <= moviesHeaderStartIndex {
		return nil, fmt.Errorf("could not find start of movie data")
	}

	entries := make([]dto.GSheetsMoviesEntry, 0)
	for _, row := range sheetData[moviesHeaderStartIndex:] {
		if row[dto.ColTMDBLinkIndex] == "" {
			// movies without a TMDB link aren't valid
			continue
		}
		entry := dto.GSheetsMoviesEntry{}
		entry.FromRowData(row)
		entries = append(entries, entry)
	}

	g.logger.Debug("Parsed", "data", entries)
	return entries, nil
}

func (g *GoogleSheetsAPI) GetMovieData() ([]dto.GSheetsMoviesEntry, error) {
	if len(g.rawGoogleDoc.Sheets) == 0 {
		g.logger.Error("no sheet found for provided spreadsheet id", "spreadSheetID", g.spreadsheetID)
		return nil, nil
	}

	mainSheetName := g.rawGoogleDoc.Sheets[0].Properties.Title
	readRange := fmt.Sprintf("%s!A:Z", mainSheetName)

	rawSheetsData, err := g.sheetsService.Spreadsheets.Values.Get(g.spreadsheetID, readRange).Do()
	if err != nil {
		g.logger.Error("could not fetch movie data from remote google sheets", "error", err)
		return nil, nil
	}

	movies, err := g.extractMoviesFromSheetData(rawSheetsData.Values)
	if err != nil {
		g.logger.Error("could not extract movie data from raw google sheets data", "error", err)
		return nil, err
	}

	return movies, nil
}
