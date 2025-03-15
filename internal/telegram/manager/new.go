package manager

import (
	"github.com/nlypage/intele"
	tele "gopkg.in/telebot.v3"
)

var Manager *intele.InputManager

func New(bot *tele.Bot) {
	Manager = intele.NewInputManager(intele.InputOptions{})
	bot.Handle(tele.OnText, Manager.MessageHandler())
	bot.Handle(tele.OnCallback, Manager.CallbackHandler())
}
