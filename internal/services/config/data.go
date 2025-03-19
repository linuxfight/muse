package config

type Group struct {
	Name          string `json:"name"`
	PlaylistId    string `json:"playlistId"`
	SheetListName string `json:"sheetListName"`
}

type Data struct {
	Bot struct {
		Token   string `json:"token"`
		Webhook *struct {
			Url    string `json:"url"`
			Secret string `json:"secret"`
		} `json:"webhook"`
		Admins []int64 `json:"admins"`
	} `json:"bot"`
	Db struct {
		Redis string `json:"redis"`
		Sheet string `json:"sheet"`
	} `json:"db"`
	Yandex struct {
		Token  string `json:"token"`
		UserId int    `json:"userId"`
	} `json:"yandex"`
	TracksLimit int     `json:"tracksLimit"`
	Groups      []Group `json:"groups"`
}
