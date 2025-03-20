package music

import (
	"time"

	"pkg.botr.me/yamusic"
)

type getPlaylistResult struct {
	InvocationInfo yamusic.InvocationInfo `json:"invocationInfo"`
	Error          error                  `json:"error"`
	Result         struct {
		PlaylistId string `json:"playlistUuid"`
		yamusic.PlaylistsResult
	} `json:"result"`
}

type playlistWithTracksResult struct {
	InvocationInfo yamusic.InvocationInfo `json:"invocationInfo"`
	Error          error                  `json:"error"`
	Result         struct {
		PlaylistId string `json:"playlistUuid"`
		yamusic.PlaylistsResult
		Tracks []struct {
			ID            int       `json:"id"`
			OriginalIndex int       `json:"originalIndex"`
			Timestamp     time.Time `json:"timestamp"`
			Track         struct {
				ID             string `json:"id"`
				RealID         string `json:"realId"`
				Title          string `json:"title"`
				ContentWarning string `json:"contentWarning"`
				Major          struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"major"`
				Available                      bool     `json:"available"`
				AvailableForPremiumUsers       bool     `json:"availableForPremiumUsers"`
				AvailableFullWithoutPermission bool     `json:"availableFullWithoutPermission"`
				AvailableForOptions            []string `json:"availableForOptions"`
				Disclaimers                    []string `json:"disclaimers"`
				StorageDir                     string   `json:"storageDir"`
				DurationMs                     int      `json:"durationMs"`
				FileSize                       int      `json:"fileSize"`
				R128                           struct {
					I  float64 `json:"i"`
					Tp float64 `json:"tp"`
				} `json:"r128"`
				Fade struct {
					InStart  float64 `json:"inStart"`
					InStop   float64 `json:"inStop"`
					OutStart float64 `json:"outStart"`
					OutStop  float64 `json:"outStop"`
				} `json:"fade"`
				PreviewDurationMs int `json:"previewDurationMs"`
				Artists           []struct {
					ID        int    `json:"id"`
					Name      string `json:"name"`
					Various   bool   `json:"various"`
					Composer  bool   `json:"composer"`
					Available bool   `json:"available"`
					Cover     struct {
						Type   string `json:"type"`
						URI    string `json:"uri"`
						Prefix string `json:"prefix"`
					} `json:"cover"`
					Genres      []interface{} `json:"genres"`
					Disclaimers []interface{} `json:"disclaimers"`
				} `json:"artists"`
				Albums []struct {
					ID             int       `json:"id"`
					Title          string    `json:"title"`
					MetaType       string    `json:"metaType"`
					ContentWarning string    `json:"contentWarning"`
					Year           int       `json:"year"`
					ReleaseDate    time.Time `json:"releaseDate"`
					CoverURI       string    `json:"coverUri"`
					Cover          struct {
						URI string `json:"uri"`
					} `json:"cover"`
					OgImage       string `json:"ogImage"`
					Genre         string `json:"genre"`
					TrackCount    int    `json:"trackCount"`
					LikesCount    int    `json:"likesCount"`
					Recent        bool   `json:"recent"`
					VeryImportant bool   `json:"veryImportant"`
					Artists       []struct {
						ID        int    `json:"id"`
						Name      string `json:"name"`
						Various   bool   `json:"various"`
						Composer  bool   `json:"composer"`
						Available bool   `json:"available"`
						Cover     struct {
							Type   string `json:"type"`
							URI    string `json:"uri"`
							Prefix string `json:"prefix"`
						} `json:"cover"`
						Genres      []interface{} `json:"genres"`
						Disclaimers []interface{} `json:"disclaimers"`
					} `json:"artists"`
					Labels []struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"labels"`
					Available                bool     `json:"available"`
					AvailableForPremiumUsers bool     `json:"availableForPremiumUsers"`
					AvailableForOptions      []string `json:"availableForOptions"`
					AvailableForMobile       bool     `json:"availableForMobile"`
					AvailablePartially       bool     `json:"availablePartially"`
					Bests                    []int    `json:"bests"`
					Disclaimers              []string `json:"disclaimers"`
					HasTrailer               bool     `json:"hasTrailer"`
					TrackPosition            struct {
						Volume int `json:"volume"`
						Index  int `json:"index"`
					} `json:"trackPosition"`
				} `json:"albums"`
				CoverURI      string `json:"coverUri"`
				DerivedColors struct {
					Average    string `json:"average"`
					WaveText   string `json:"waveText"`
					MiniPlayer string `json:"miniPlayer"`
					Accent     string `json:"accent"`
				} `json:"derivedColors"`
				OgImage            string `json:"ogImage"`
				LyricsAvailable    bool   `json:"lyricsAvailable"`
				Type               string `json:"type"`
				RememberPosition   bool   `json:"rememberPosition"`
				BackgroundVideoURI string `json:"backgroundVideoUri"`
				TrackSharingFlag   string `json:"trackSharingFlag"`
				PlayerID           string `json:"playerId"`
				LyricsInfo         struct {
					HasAvailableSyncLyrics bool `json:"hasAvailableSyncLyrics"`
					HasAvailableTextLyrics bool `json:"hasAvailableTextLyrics"`
				} `json:"lyricsInfo"`
				TrackSource           string   `json:"trackSource"`
				SpecialAudioResources []string `json:"specialAudioResources"`
			} `json:"track"`
			Recent               bool `json:"recent"`
			OriginalShuffleIndex int  `json:"originalShuffleIndex"`
		} `json:"tracks"`
	} `json:"result"`
}

type listResponse struct {
	InvocationInfo yamusic.InvocationInfo `json:"invocationInfo"`
	Error          error                  `json:"error"`
	Result         []struct {
		PlaylistId string `json:"playlistUuid"`
		yamusic.PlaylistsResult
	} `json:"result"`
}
