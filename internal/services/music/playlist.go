package music

import (
	"context"
	"encoding/json"
	"io"
	"muse/internal/errorz"
	"muse/internal/services/logger"

	"pkg.botr.me/yamusic"
)

type playlist struct {
	PlaylistId string                   `json:"playlistUuid"`
	Tracks     []yamusic.PlaylistsTrack `json:"tracks"`
	yamusic.PlaylistsResult
}

type listResponse struct {
	InvocationInfo yamusic.InvocationInfo `json:"invocationInfo"`
	Error          error                  `json:"error"`
	Result         []playlist             `json:"result"`
}

func (s *Service) GeneratePlaylist(c context.Context, playlistId string, tracks *[]yamusic.PlaylistsTrack) error {
	_, resp, err := s.client.Playlists().List(c, s.client.UserID())
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log.Panicf("Failed to close body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Panicf("failed to read body: %v", err)
		return err
	}
	list := listResponse{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		logger.Log.Panicf("Failed to unmarshal response: %v", err)
		return err
	}

	for _, playlist := range list.Result {
		if playlist.PlaylistId == playlistId {
			s.client.Playlists().RemoveTracks(c, playlist.Kind, playlist.Revision, playlist.Tracks, nil)
			_, _, err = s.client.Playlists().AddTracks(c, playlist.Kind, playlist.Revision, *tracks, nil)
			if err != nil {
				return err
			}
		}
	}

	return errorz.PlaylistNotFound
}
