package music

import (
	"muse/internal/services/logger"
	"os"
	"pkg.botr.me/yamusic"
	"strconv"
)

type Service struct {
	client *yamusic.Client
}

func New() *Service {
	// Yandex Music API init
	ymToken := os.Getenv("YANDEX_ACCESS_TOKEN")
	if ymToken == "" {
		logger.Log.Fatal("YANDEX_ACCESS_TOKEN environment variable not set")
	}
	ymId, err := strconv.Atoi(os.Getenv("YANDEX_USER_ID"))
	if err != nil {
		logger.Log.Fatalf("YANDEX_USER_ID environment variable not set, or set incorrectly: %v", err)
	}

	return &Service{
		client: yamusic.NewClient(
			yamusic.AccessToken(ymId, ymToken),
		),
	}
}
