package errorz

import "errors"

var (
	TrackNotFound      = errors.New("трек не найден")
	TrackExists        = errors.New("трек уже добавлен")
	NoTracks           = errors.New("нет треков в плейлисте")
	TracksLimitReached = errors.New("лимит треков в плейлисте превышен")
	PlaylistNotFound   = errors.New("плейлист не найден")
)
