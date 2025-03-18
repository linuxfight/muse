package config

import (
	"encoding/json"
	"muse/internal/services/logger"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	AdminIds    []int64
	Groups      []Group
	TracksLimit int
}

type Group struct {
	Name          string `json:"name"`
	PlaylistId    string `json:"playlistId"`
	SheetListName string `json:"sheetListName"`
}

func New() *Data {
	groups := []Group{}

	data, err := os.ReadFile("settings/playlists.json")
	if err != nil {
		logger.Log.Panicf("failed to read settings/playlists.json: %v", err)
		return nil
	}

	if err := json.Unmarshal(data, &groups); err != nil {
		logger.Log.Panicf("failed to unmarshal settings/playlists.json: %v", err)
		return nil
	}

	if len(groups) == 0 {
		logger.Log.Panicf("no groups found in settings/playlists.json")
		return nil
	}

	adminString := os.Getenv("BOT_ADMINS")
	adminStringSlice := strings.Split(adminString, ",")
	var adminIds []int64
	for _, adminString := range adminStringSlice {
		adminId, err := strconv.ParseInt(adminString, 10, 64)
		if err != nil {
			logger.Log.Fatalf("Failed to parse admin id from %s: %v", adminString, err)
		}
		adminIds = append(adminIds, adminId)
	}

	tracksLimitString := os.Getenv("TRACKS_LIMIT")
	tracksLimit, err := strconv.Atoi(tracksLimitString)
	if err != nil {
		logger.Log.Fatalf("Failed to parse tracks limit from %s: %v", tracksLimitString, err)
		return nil
	}

	return &Data{
		AdminIds:    adminIds,
		Groups:      groups,
		TracksLimit: tracksLimit,
	}
}

func (d *Data) GetGroup(playlistId string) *Group {
	for _, group := range d.Groups {
		if group.PlaylistId == playlistId {
			return &group
		}
	}
	return nil
}
