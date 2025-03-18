package handlers

import (
	"muse/internal/services/config"
	"slices"

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

func groupMenu(groups []config.Group) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	rows := []tele.Row{}
	for _, group := range groups {
		rows = append(rows, menu.Row(tele.Btn{
			Text:   group.Name,
			Unique: group.PlaylistId,
		}))
	}

	menu.Inline(rows...)

	return menu
}
