package handlers

import (
	tele "gopkg.in/telebot.v3"
)

func (ctl *Controller) greet(c tele.Context) error {
	return c.Send("Привет, я бот для плейлиста будущей дискотеки! 🎧\n \n"+
		"Мы заботимся о том, чтобы ваша дискотека была проведена в наилучшем виде ❤\n \n"+
		"Отправь мне песню и я добавлю её в плейлист, если она пройдёт модерацию, то она будет в плейлисте ✨", markup(c.Sender().ID, ctl.admins))
}
