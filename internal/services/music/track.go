package music

import (
	"context"
	"muse/internal/errorz"
	"pkg.botr.me/yamusic"
)

func (s *Service) GetTrack(c context.Context, query string) (*yamusic.Track, error) {
	tracks, _, err := s.client.Search().Tracks(c, query, nil)
	if err != nil {
		return nil, err
	}

	if len(tracks.Result.Tracks.Results) < 1 {
		return nil, errorz.TrackNotFound
	}

	get, _, err := s.client.Tracks().Get(c, tracks.Result.Tracks.Results[0].ID)
	if err != nil {
		return nil, err
	}

	if len(get.Result) < 1 {
		return nil, errorz.TrackNotFound
	}

	return &get.Result[0], nil
}
