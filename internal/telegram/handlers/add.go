package handlers

import (
	"context"
	"fmt"
	"muse/internal/errorz"
	"muse/internal/services/logger"
	"muse/internal/telegram/manager"
	"strings"

	"github.com/nlypage/intele/collector"
	tele "gopkg.in/telebot.v3"
)

func (ctl *Controller) addNew(ctx tele.Context) error {
	if err := ctx.Respond(); err != nil {
		logger.Log.Errorf("failed to respond: %v", err)
		return err
	}

	playlistId, tracksCount, err := ctl.storage.GetUser(context.Background(), ctx.Sender().ID)

	if err != nil {
		return ctl.greet(ctx)
	}

	if tracksCount >= ctl.config.TracksLimit {
		return ctx.Send(fmt.Sprintf("Ошибка: %s", errorz.TracksLimitReached))
	}

	group := ctl.config.GetGroup(playlistId)
	if group == nil {
		return ctl.greet(ctx)
	}

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
		return nil
	}

	track, err := ctl.music.GetTrack(context.Background(), *stage.Value)
	if err != nil {
		logger.Log.Error(err)
		return ctx.Send(fmt.Sprintf("Ошибка: %s", err.Error()))
	}

	exists, err := ctl.sheets.Exists(track.ID, group.SheetListName)
	if err != nil {
		logger.Log.Error(err)
		return ctx.Send(fmt.Sprintf("Ошибка: %s", err.Error()))
	}

	if exists {
		return ctx.Send(fmt.Sprintf("Ошибка: %s", errorz.TrackExists))
	}

	if err := ctl.sheets.Insert(track, group.SheetListName); err != nil {
		if err := ctx.Send(notFoundText); err != nil {
			logger.Log.Error(err)
			return ctx.Send(fmt.Sprintf("Ошибка: %s", err.Error()))
		}
	}

	if err := ctl.storage.UpdateUser(context.Background(), ctx.Sender().ID, playlistId, tracksCount+1); err != nil {
		logger.Log.Error(err)
		return ctx.Send(fmt.Sprintf("Ошибка: %s", err.Error()))
	}

	artists := ""
	for _, artist := range track.Artists {
		artists += artist.Name + ", "
	}
	artists = strings.TrimSuffix(artists, ", ")

	logger.Log.Infof("%d added %s - %s", ctx.Sender().ID, artists, track.Title)

	return ctx.Send(fmt.Sprintf("Добавлен трек: %s - %s", artists, track.Title))
}
