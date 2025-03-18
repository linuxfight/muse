package sheets

import (
	"fmt"
	"strings"

	"google.golang.org/api/sheets/v4"
	"pkg.botr.me/yamusic"
)

func (s *Service) Insert(track *yamusic.Track, sheetListName string) error {
	artists := ""
	for _, artist := range track.Artists {
		artists += artist.Name + ", "
	}
	artists = strings.TrimSuffix(artists, ", ")

	link := fmt.Sprintf("https://music.yandex.ru/album/%d/track/%s", track.Albums[0].ID, track.ID)

	duration := track.DurationMs / 1000 / 60

	values := &sheets.ValueRange{
		Values: [][]interface{}{
			// link - https://music.yandex.ru/album/5543883/track/42083820
			// ID, AlbumId, Artist, Title, Link, Explicit, Time, Add (net)
			{track.ID, track.Albums[0].ID, artists, track.Title, track.ContentWarning == "explicit",
				link, duration, false},
		},
	}

	rangeToAppend := fmt.Sprintf("%s!A:F", sheetListName)

	_, err := s.sheetsClient.Spreadsheets.Values.Append(
		s.sheetId,
		rangeToAppend,
		values,
	).ValueInputOption("USER_ENTERED").Do()

	return err
}
