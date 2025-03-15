package handlers

import (
	tele "gopkg.in/telebot.v3"
	"slices"
)

var (
	addNewBtn = tele.Btn{
		Text:   "Добавить трек",
		Unique: "add_new",
	}

	genPlaylistBtn = tele.Btn{
		Text:   "Создать плейлист",
		Unique: "gen_playlist",
	}
)

func markup(userId int64, admins []int64) *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	if slices.Contains(admins, userId) {
		menu.Inline(menu.Row(addNewBtn), menu.Row(genPlaylistBtn))
	} else {
		menu.Inline(menu.Row(addNewBtn))
	}

	return menu
}
