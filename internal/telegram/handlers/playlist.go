package handlers

import (
	"context"
	"muse/internal/services/logger"

	tele "gopkg.in/telebot.v3"
)

func (ctl *Controller) generatePlaylist(ctx tele.Context) error {
	if err := ctx.Respond(); err != nil {
		logger.Log.Errorf("failed to respond: %v", err)
		return err
	}

	playlistId, _, err := ctl.storage.GetUser(context.Background(), ctx.Sender().ID)
	if err != nil {
		return ctl.greet(ctx)
	}

	group := ctl.config.GetGroup(playlistId)
	if group == nil {
		return ctl.greet(ctx)
	}

	tracks, err := ctl.sheets.GetAllTracks(group.SheetListName)
	if err != nil {
		logger.Log.Errorf("failed to get all tracks: %v", err)
		return ctx.Send("Ошибка при обновлении плейлиста, обратитесь к администратору")
	}

	if err := ctx.Send("Треки получены, запущено обновление плейлиста"); err != nil {
		logger.Log.Errorf("failed to get all tracks: %v", err)
		return ctx.Send("Ошибка при обновлении плейлиста, обратитесь к администратору")
	}

	if err := ctl.music.GeneratePlaylist(context.Background(), playlistId, tracks); err != nil {
		logger.Log.Errorf("failed to generate playlist: %v", err)
		return ctx.Send("Ошибка при обновлении плейлиста, обратитесь к администратору")
	}

	return ctx.Send("Обновление завершено!")
}
