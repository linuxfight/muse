package handlers

import (
	"muse/internal/services/config"
	"slices"

	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

var (
	addNewBtn = tele.Btn{
		Text:   "Добавить трек",
		Unique: "add_new",
	}

	genPlaylistBtn = tele.Btn{
		Text:   "Обновить плейлисты",
		Unique: "gen_playlist",
	}
)

func startMenu(userId int64, admins []int64) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	if slices.Contains(admins, userId) {
		menu.Inline(menu.Row(addNewBtn), menu.Row(genPlaylistBtn))
	} else {
		menu.Inline(menu.Row(addNewBtn))
	}

	return menu
}

func groupMenu(groups []config.Group) []telebot.Btn {
	rows := []tele.Btn{}
	for _, group := range groups {
		rows = append(rows, tele.Btn{
			Text:   group.Name,
			Unique: "group",
			Data:   group.PlaylistId,
		})
	}

	return rows
}
