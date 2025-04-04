package handlers

import (
	"context"
	"fmt"
	"muse/internal/services/logger"
	"muse/internal/telegram/manager"

	"github.com/nlypage/intele/collector"
	tele "gopkg.in/telebot.v3"
)

const text = "Привет, я бот для плейлиста будущей дискотеки! 🎧\n \n" +
	"Мы заботимся о том, чтобы ваша дискотека была проведена в наилучшем виде ❤\n \n" +
	"Отправь мне песню и я добавлю её в плейлист, если она пройдёт модерацию, то она будет в плейлисте ✨"

func (ctl *Controller) greet(ctx tele.Context) error {
	_, tracksCount, err := ctl.storage.GetUser(context.Background(), ctx.Sender().ID)
	if err != nil {
		if err := ctx.Send(text); err != nil {
			logger.Log.Errorf("failed to respond: %v", err)
			return err
		}

		playlistId := ctl.config.Groups[0].PlaylistId

		if len(ctl.config.Groups) >= 2 {
			inputCollector := collector.New()
			stage := &manager.InputStage{
				Prompt:       "Выберете свою группу для плейлиста",
				ErrorMessage: "Ошибка, попробуйте ещё раз",
				IsValid: func(s string) bool {
					return false
				},
				Value: nil,
			}

			if err := manager.PromptForInput(inputCollector, ctx, stage, groupMenu(ctl.config.Groups)...); err != nil {
				logger.Log.Errorf("failed getting start input: %v", err)
				return ctx.Send(fmt.Sprintf("Ошибка: %s", err.Error()))
			}

			if *stage.Value == "" {
				logger.Log.Errorf("start input is empty")
				return ctx.Send("Введите /start, чтобы начать регистрацию")
			}

			playlistId = *stage.Value
		}

		if err := ctl.storage.UpdateUser(context.Background(), ctx.Sender().ID, playlistId, 0); err != nil {
			logger.Log.Errorf("failed to update redis: %s", err.Error())
			return ctx.Send(fmt.Sprintf("Ошибка: %s", err.Error()))
		}
	}

	return ctx.Send(fmt.Sprintf("❗ Ваш лимит треков: %d. \n\n❔ Что вы хотите сделать?", ctl.config.TracksLimit-tracksCount), startMenu(ctx.Sender().ID, ctl.config.Bot.Admins))
}
