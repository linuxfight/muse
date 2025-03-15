package sheets

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"muse/internal/services/logger"
	"net/http"
	"os"
)

type Service struct {
	sheetsClient *sheets.Service
	sheetId      string
}

// New Google Sheets API init
func New(ctx context.Context) *Service {
	sheetsClient, err := getClient()
	if err != nil {
		logger.Log.Fatalf("Failed to initialize Sheets client: %v", err)
	}
	sheetId := os.Getenv("SHEET_ID")
	if sheetId == "" {
		logger.Log.Fatalf("SHEET_ID environment variable not set")
	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(sheetsClient))
	if err != nil {
		logger.Log.Fatalf("Failed to initialize Sheets client: %v", err)
	}

	return &Service{
		sheetsClient: srv,
		sheetId:      sheetId,
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient() (*http.Client, error) {
	b, err := os.ReadFile("settings/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tokFile := "settings/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		return nil, err
	}
	return config.Client(context.Background(), tok), nil
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logger.Log.Panicf("Unable to close file: %v", err)
		}
	}(f)
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
