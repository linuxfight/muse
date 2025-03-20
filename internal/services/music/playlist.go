package music

import (
	"context"
	"encoding/json"
	"io"
	"muse/internal/errorz"
	"muse/internal/services/logger"
	"strconv"

	"pkg.botr.me/yamusic"
)

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
	if err := json.Unmarshal(body, &list); err != nil {
		logger.Log.Panicf("Failed to unmarshal response: %v", err)
		return err
	}

	for _, p := range list.Result {
		if p.PlaylistId == playlistId {
			_, resp, err := s.client.Playlists().Get(c, s.client.UserID(), p.Kind)
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
			playlist := playlistWithTracksResult{}
			if err := json.Unmarshal(body, &playlist); err != nil {
				logger.Log.Panicf("Failed to unmarshal response: %v", err)
				return err
			}

			tracksToRemove := []yamusic.PlaylistsTrack{}
			for _, track := range playlist.Result.Tracks {
				id, err := strconv.Atoi(track.Track.ID)
				if err != nil {
					logger.Log.Panicf("failed to parse int: %s", track.ID)
					return err
				}

				tracksToRemove = append(tracksToRemove, yamusic.PlaylistsTrack{
					AlbumID: track.Track.Albums[0].ID,
					ID:      id,
				})
			}

			revision := playlist.Result.Revision

			if len(tracksToRemove) > 0 {
				if _, _, err = s.client.Playlists().RemoveTracks(c, playlist.Result.Kind, revision, tracksToRemove, nil); err != nil {
					return err
				}
				revision += 1
			}

			if _, _, err = s.client.Playlists().AddTracks(c, playlist.Result.Kind, revision, *tracks, nil); err != nil {
				return err
			}

			return nil
		}
	}

	return errorz.PlaylistNotFound
}
