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

	_, _, err := ctl.storage.GetUser(context.Background(), ctx.Sender().ID)
	if err != nil {
		return ctl.greet(ctx)
	}

	for _, group := range ctl.config.Groups {
		tracks, err := ctl.sheets.GetAllTracks(group.SheetListName)
		if err != nil {
			logger.Log.Errorf("failed to get all tracks: %v", err)
			return ctx.Send("Ошибка при обновлении плейлиста %s, обратитесь к администратору", group.PlaylistId)
		}

		if err := ctx.Send("Треки получены, запущено обновление плейлиста"); err != nil {
			logger.Log.Errorf("failed to get all tracks: %v", err)
			return ctx.Send("Ошибка при обновлении плейлиста s, обратитесь к администратору", group.PlaylistId)
		}

		if err := ctl.music.GeneratePlaylist(context.Background(), group.PlaylistId, tracks); err != nil {
			logger.Log.Errorf("failed to generate playlist: %v", err)
			return ctx.Send("Ошибка при обновлении плейлиста, обратитесь к администратору")
		}
	}

	return ctx.Send("Обновление завершено!")
}
