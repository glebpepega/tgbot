package sender

import (
	"github.com/glebpepega/goodvibesbot/sendable/photo"
	"github.com/glebpepega/goodvibesbot/sendable/text"
	"github.com/glebpepega/goodvibesbot/update"
)

func Sender(ch chan update.Update) {
	go func() {
		for update := range ch {
			if update.Message.Text == "/get" {
				SendPic(update.Message.Chat.ID)
			} else {
				SendText(update.Message.Chat.ID, "Click /get")
			}
			if update.Callback_Query.ID != "" {
				update.Callback_Query.Answer()
				SendPic(update.Callback_Query.From.ID)
			}
		}
	}()
}

func SendPic(id int) {
	ph := photo.New()
	ph.Send(id)
}

func SendText(id int, message string) {
	tx := text.New()
	tx.Send(id, message)
}
