package handlers

import (
	"context"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"muse/internal/services/logger"
	"muse/internal/services/music"
	"muse/internal/services/sheets"
	"muse/internal/telegram/manager"
	"os"
	"time"
)

type Controller struct {
	bot     *tele.Bot
	music   *music.Service
	sheets  *sheets.Service
	admins  []int64
	dev     bool
	webhook string
}

func New() *Controller {
	// init settings
	// TODO: get ids from env
	adminIds := []int64{
		687627953,
	}
	debug := os.Getenv("DEBUG") != ""

	// Logger init
	logger.New(debug)

	// Telegram Bot init
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		logger.Log.Fatal("BOT_TOKEN environment variable not set")
	}

	var poller tele.Poller
	if debug {
		poller = &tele.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "callback_query"},
		}
	} else {
		poller = &tele.Webhook{
			Listen:         "0.0.0.0:8080",
			MaxConnections: 30,
			AllowedUpdates: []string{"message", "callback_query"},
			IP:             os.Getenv("WEBHOOK_IP"),
			DropUpdates:    true,
			SecretToken:    os.Getenv("WEBHOOK_SECRET"),
			HasCustomCert:  false,
			TLS:            nil,
			Endpoint: &tele.WebhookEndpoint{
				PublicURL: os.Getenv("WEBHOOK_URL"),
			},
		}
	}

	pref := tele.Settings{
		Token:  botToken,
		Poller: poller,
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		logger.Log.Fatal(err)
		return nil
	}

	// set up middlewares
	if debug {
		bot.Use(middleware.Logger()) // TODO: switch to opentelemetry logging and tracing, enable only in DEV
	} else {
		bot.Use(middleware.Recover())
	}

	ctl := &Controller{
		bot:    bot,
		music:  music.New(),
		sheets: sheets.New(context.Background()),
		admins: adminIds,
		dev:    debug,
	}

	// set up user handlers
	manager.New(bot)
	bot.Handle("/start", func(ctx tele.Context) error { return ctl.greet(ctx) })
	bot.Handle(&addNewBtn, func(ctx tele.Context) error { return ctl.addNew(ctx) })

	// set up admin handlers
	admin := bot.Group()
	admin.Use(middleware.Whitelist(adminIds...))
	admin.Handle(&genPlaylistBtn, func(ctx tele.Context) error { return ctl.generatePlaylist(ctx) })

	return ctl
}

func (ctl *Controller) Start() {
	logger.Log.Infof("Logged in as: https://t.me/@%s", ctl.bot.Me.Username)
	err := ctl.bot.RemoveWebhook(true)
	if err != nil {
		logger.Log.Errorf("Failed to remove webhook: %v", err)
		return
	}
	ctl.bot.Start()
}
