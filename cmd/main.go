package main

import (
	"muse/internal/handlers"
	"muse/internal/logger"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	logger.New(false, "")

	b, err := tele.NewBot(pref)
	if err != nil {
		logger.Log.Fatal(err)
		return
	}

	logger.Log.Infof("Logged in as: https://t.me/@%s", b.Me.Username)

	b.Handle("/start", func(c tele.Context) error {
		return handlers.Start(c)
	})

	logger.Log.Info("Bot started")
	b.Start()
}
