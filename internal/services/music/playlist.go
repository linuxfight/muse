package music

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"muse/internal/services/logger"
	"net/http"
	"net/url"
	"pkg.botr.me/yamusic"
	"strconv"
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

	_, _, err = s.addTracks(c, create.Result, *tracks, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://music.yandex.ru/playlists/%s", create.Result.PlaylistId), nil
}

// adds tracks to playlist
func (s *Service) addTracks(
	ctx context.Context,
	source playlist,
	tracks []yamusic.PlaylistsTrack,
	opts *yamusic.PlaylistsAddTracksOptions,
) (*yamusic.PlaylistsAddTracksResp, *http.Response, error) {
	if opts == nil {
		opts = &yamusic.PlaylistsAddTracksOptions{
			At: 0,
		}
	}

	diff := []struct {
		Op     string                   `json:"op"`
		At     int                      `json:"at"`
		Tracks []yamusic.PlaylistsTrack `json:"tracks"`
	}{
		{
			Op:     "insert",
			At:     opts.At,
			Tracks: tracks,
		},
	}

	b, err := json.Marshal(diff)
	if err != nil {
		return nil, nil, err
	}

	form := url.Values{}
	form.Set("diff", string(b))
	form.Set("revision", strconv.Itoa(source.Revision))

	uri := fmt.Sprintf(
		"users/%v/playlists/%v/change-relative",
		source.Owner.UID,
		source.Kind,
	)

	req, err := s.client.NewRequest(http.MethodPost, uri, form)
	if err != nil {
		return nil, nil, err
	}

	addTracksResp := new(yamusic.PlaylistsAddTracksResp)
	resp, err := s.client.Do(ctx, req, addTracksResp)
	return addTracksResp, resp, err
}
