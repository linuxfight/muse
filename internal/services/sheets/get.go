package sheets

import (
	"fmt"
	"muse/internal/errorz"
	"pkg.botr.me/yamusic"
	"strconv"
)

func (s *Service) Exists(id string) (bool, error) {
	rangeToRead := fmt.Sprintf("%s!A2:A", sheetName)
	resp, err := s.sheetsClient.Spreadsheets.Values.Get(s.sheetId, rangeToRead).Do()
	if err != nil {
		return false, err
	}

	if len(resp.Values) == 0 {
		return false, nil
	}

	for _, row := range resp.Values {
		if len(row) > 0 {
			cellValue := fmt.Sprintf("%v", row[0])
			if cellValue == id {
				return true, nil
			}
		}
	}

	return false, nil
}

func (s *Service) GetAllTracks() (*[]yamusic.PlaylistsTrack, error) {
	rangeToRead := fmt.Sprintf("%s!A2:H", sheetName)
	resp, err := s.sheetsClient.Spreadsheets.Values.Get(s.sheetId, rangeToRead).Do()
	if err != nil {
		return nil, err
	}

	if len(resp.Values) == 0 {
		return nil, errorz.NoTracks
	}

	tracks := make([]yamusic.PlaylistsTrack, len(resp.Values))
	for i, row := range resp.Values {
		id, err := strconv.Atoi(row[0].(string))
		if err != nil {
			return nil, err
		}
		albumId, err := strconv.Atoi(row[1].(string))
		if err != nil {
			return nil, err
		}
		allowed := row[7]

		if allowed == "FALSE" {
			continue
		}

		tracks[i] = yamusic.PlaylistsTrack{
			ID:      id,
			AlbumID: albumId,
		}
	}

	return &tracks, nil
}
