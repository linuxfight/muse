package webhook

import (
	"context"
	"github.com/gofiber/fiber/v3"
	tele "gopkg.in/telebot.v3"
	"muse/internal/services/logger"
)

// A Poller configures the poller for webhooks. It opens a port on the given
// listen address. If TLS is filled, the listener will use the key and cert to open
// a secure port. Otherwise it will use plain HTTP.
//
// If you have a loadbalancer ore other infrastructure in front of your service, you
// must fill the Endpoint structure so this poller will send this data to telegram. If
// you leave these values empty, your local address will be sent to telegram which is mostly
// not what you want (at least while developing). If you have a single instance of your
// bot you should consider to use the LongPoller instead of a WebHook.
//
// You can also leave the Listen field empty. In this case it is up to the caller to
// add the Webhook to a http-mux.
type Poller struct {
	webhook *tele.Webhook

	dest chan<- tele.Update
	bot  *tele.Bot

	listen      string
	secretToken string
}

func New(webhook *tele.Webhook) *Poller {
	return &Poller{
		webhook:     webhook,
		listen:      webhook.Listen,
		secretToken: webhook.SecretToken,
	}
}

// Poll starts the Fiber server to listen for webhook updates.
func (h *Poller) Poll(b *tele.Bot, dest chan tele.Update, stop chan struct{}) {
	if err := b.SetWebhook(h.webhook); err != nil {
		b.OnError(err, nil)
		close(stop)
		return
	}

	h.dest = dest
	h.bot = b

	if h.listen == "" {
		h.waitForStop(stop)
		return
	}

	app := fiber.New()
	app.Post("/", h.fiberHandler)

	shutdown := make(chan struct{})
	go func() {
		<-stop
		err := app.ShutdownWithContext(context.Background())
		if err != nil {
			return
		}
		close(shutdown)
	}()

	if err := app.Listen(h.listen); err != nil {
		b.OnError(err, nil)
	}

	<-shutdown
	close(stop)
}

// fiberHandler processes incoming webhook requests using Fiber.
func (h *Poller) fiberHandler(c fiber.Ctx) error {
	if h.secretToken != "" && c.Get("X-Telegram-Bot-Api-Secret-Token") != h.secretToken {
		logger.Log.Debugf("invalid secret token in request")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var update tele.Update
	if err := c.Bind().JSON(&update); err != nil {
		logger.Log.Debugf("failed to bind update: %v", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	h.dest <- update
	return c.SendStatus(fiber.StatusOK)
}

// waitForStop handles stopping when no server is started.
func (h *Poller) waitForStop(stop chan struct{}) {
	<-stop
	close(stop)
}
