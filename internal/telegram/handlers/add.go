package handlers

import (
	"context"
	"fmt"
	"github.com/nlypage/intele/collector"
	tele "gopkg.in/telebot.v3"
	"muse/internal/errorz"
	"muse/internal/services/logger"
	"muse/internal/telegram/manager"
	"strings"
)

func (ctl *Controller) addNew(ctx tele.Context) error {
	inputCollector := collector.New()
	notFoundText := fmt.Sprintf("Ошибка: %s", errorz.TrackNotFound.Error())

	stage := &manager.InputStage{
		Prompt:       "Напиши НАЗВАНИЕ песни или ИСПОЛНИТЕЛЬ - НАЗВАНИЕ, вот так: Memory Reboot или VØJ - Memory Reboot",
		ErrorMessage: notFoundText,
		IsValid: func(s string) bool {
			return len(s) > 0
		},
		Value: nil,
	}

	if err := manager.PromptForInput(inputCollector, ctx, stage); err != nil {
		return err
	}

	if stage.Value == nil {
		return ctx.Respond()
	}

	track, err := ctl.music.GetTrack(context.Background(), *stage.Value)
	if err != nil {
		return err
	}

	exists, err := ctl.sheets.Exists(track.ID)
	if err != nil {
		logger.Log.Error(err)
		return err
	}

	if exists {
		return ctx.RespondText(fmt.Sprintf("Ошибка: %s", errorz.TrackExists))
	}

	if err := ctl.sheets.Insert(track); err != nil {
		if err := ctx.RespondText(notFoundText); err != nil {
			logger.Log.Error(err)
			return err
		}
	}

	artists := ""
	for _, artist := range track.Artists {
		artists += artist.Name + ", "
	}
	artists = strings.TrimSuffix(artists, ", ")

	logger.Log.Infof("%d added %s - %s", ctx.Sender().ID, artists, track.Title)

	return ctx.RespondText(fmt.Sprintf("Добавлен трек: %s - %s", artists, track.Title))
}
