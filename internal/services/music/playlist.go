package music

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"muse/internal/services/logger"
	"pkg.botr.me/yamusic"
)

type playlist struct {
	PlaylistId string `json:"playlistUuid"`
	yamusic.PlaylistsResult
}

type createResponse struct {
	InvocationInfo yamusic.InvocationInfo `json:"invocationInfo"`
	Error          error                  `json:"error"`
	Result         playlist               `json:"result"`
}

func (s *Service) GeneratePlaylist(c context.Context, title string, tracks *[]yamusic.PlaylistsTrack) (string, error) {
	_, resp, err := s.client.Playlists().Create(c, title, true)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log.Panicf("Failed to close body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	create := createResponse{}
	err = json.Unmarshal(body, &create)
	if err != nil {
		logger.Log.Panicf("Failed to unmarshal response: %v", err)
		return "", err
	}

	_, _, err = s.client.Playlists().AddTracks(c, create.Result.Kind, create.Result.Revision, *tracks, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://music.yandex.ru/playlists/%s", create.Result.PlaylistId), nil
}
