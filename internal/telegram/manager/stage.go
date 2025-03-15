package manager

import (
	"context"
	"github.com/nlypage/intele"
	"github.com/nlypage/intele/collector"
	"gopkg.in/telebot.v3"
	"time"
)

type InputStage struct {
	Prompt       string
	ErrorMessage string
	IsValid      func(string) bool
	Value        *string
}

// idea from here https://github.com/Badsnus/cu-clubs-bot/blob/main/bot/internal/adapters/controller/telegram/handlers/clubOwner/clubOwner.go

func PromptForInput(coll *collector.MessageCollector, c telebot.Context, stage *InputStage, buttons ...telebot.Btn) error {
	if buttons != nil {
		menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
		var rows []telebot.Row
		for _, btn := range buttons {
			rows = append(rows, menu.Row(btn))
		}
		menu.Inline(
			rows...,
		)
		_ = coll.Send(c, stage.Prompt, menu)
	} else {
		_ = coll.Send(c, stage.Prompt)
	}

	for {
		var response intele.Response
		var err error
		if buttons != nil {
			var bends []telebot.CallbackEndpoint
			for _, b := range buttons {
				bends = append(bends, &b)
			}
			response, err = Manager.Get(context.Background(), c.Sender().ID, 90*time.Second, bends...)
		} else {
			response, err = Manager.Get(context.Background(), c.Sender().ID, 90*time.Second)
		}
		if err == nil && response.Message != nil {
			coll.Collect(response.Message)
		}

		switch {
		case response.Canceled:
			_ = coll.Clear(c, collector.ClearOptions{
				IgnoreErrors: true,
				ExcludeLast:  true,
			})
			return nil
		case err != nil:
			_ = coll.Clear(c, collector.ClearOptions{
				IgnoreErrors: true,
			})
			return nil
		case response.Callback != nil:
			return coll.Clear(c, collector.ClearOptions{
				IgnoreErrors: true,
			})
		case stage.IsValid(response.Message.Text):
			stage.Value = &response.Message.Text
			return coll.Clear(c, collector.ClearOptions{
				IgnoreErrors: true,
			})
		default:
			_ = coll.Clear(c, collector.ClearOptions{
				IgnoreErrors: true,
			})
			_ = coll.Send(c,
				stage.ErrorMessage,
			)
		}
	}
}
