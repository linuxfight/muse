package music

import "pkg.botr.me/yamusic"

type Service struct {
	client *yamusic.Client
}

func New(ymId int, ymToken string) *Service {
	return &Service{
		client: yamusic.NewClient(
			yamusic.AccessToken(ymId, ymToken),
		),
	}
}
