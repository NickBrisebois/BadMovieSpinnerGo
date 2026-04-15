package external

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetsAPI struct {
	sheetsService *sheets.Service
}

func NewGoogleSheetsAPI(credentialsFile, spreadsheetID string) (*GoogleSheetsAPI, error) {
	sheetsClientOptions := option.WithAuthCredentialsFile(option.ServiceAccount, credentialsFile)
	sheetsService, err := sheets.NewService(
		context.Background(),
		sheetsClientOptions,
		option.WithScopes(sheets.SpreadsheetsReadonlyScope),
	)
	if err != nil {
		return nil, err
	}
	return &GoogleSheetsAPI{sheetsService: sheetsService}, nil
}
