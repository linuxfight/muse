package config

import (
	"encoding/json"
	"muse/internal/services/logger"
	"os"
)

func New() *Data {
	data := Data{}

	file, err := os.ReadFile("settings/config.json")
	if err != nil {
		logger.Log.Panicf("failed to read settings/config.json: %v", err)
		return nil
	}

	if err := json.Unmarshal(file, &data); err != nil {
		logger.Log.Panicf("failed to unmarshal settings/config.json: %v", err)
		return nil
	}

	if len(data.Groups) == 0 {
		logger.Log.Panicf("no groups found in settings/config.json")
		return nil
	}

	if len(data.Bot.Admins) == 0 {
		logger.Log.Panicf("no groups found in settings/config.json")
		return nil
	}

	return &data
}

func (d *Data) GetGroup(playlistId string) *Group {
	for _, group := range d.Groups {
		if group.PlaylistId == playlistId {
			return &group
		}
	}
	return nil
}
