package handlers

import (
	"context"
	"github.com/nlypage/intele/collector"
	tele "gopkg.in/telebot.v3"
	"muse/internal/services/logger"
	"muse/internal/telegram/manager"
)

func (ctl *Controller) generatePlaylist(ctx tele.Context) error {
	err := ctx.Respond()
	if err != nil {
		logger.Log.Errorf("failed to respond: %v", err)
		return err
	}

	tracks, err := ctl.sheets.GetAllTracks()
	if err != nil {
		logger.Log.Errorf("failed to get all tracks: %v", err)
		return ctx.Send("Ошибка при создании плейлиста, обратитесь к администратору")
	}

	inputCollector := collector.New()
	stage := &manager.InputStage{
		Prompt:       "Введите название плейлиста",
		ErrorMessage: "Ошибка, попробуйте ещё раз",
		IsValid: func(s string) bool {
			return len(s) > 0
		},
		Value: nil,
	}

	if err := manager.PromptForInput(inputCollector, ctx, stage); err != nil {
		return err
	}

	if stage.Value == nil {
		return nil
	}

	err = ctx.Send("Треки получены, запущено создание плейлиста")
	if err != nil {
		logger.Log.Errorf("failed to get all tracks: %v", err)
		return ctx.Send("Ошибка при создании плейлиста, обратитесь к администратору")
	}

	playlistLink, err := ctl.music.GeneratePlaylist(context.Background(), *stage.Value, tracks)
	if err != nil {
		logger.Log.Errorf("failed to generate playlist: %v", err)
		return ctx.Send("Ошибка при создании плейлиста, обратитесь к администратору")
	}

	return ctx.Send("Создание завершено! " + playlistLink)
}
