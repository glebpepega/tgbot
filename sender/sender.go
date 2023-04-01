package sender

import (
	"github.com/glebpepega/goodvibesbot/photo"
	"github.com/glebpepega/goodvibesbot/update"
)

func Sender(ch chan update.Update) {
	go func() {
		for update := range ch {
			if len(update.Message.Entities) > 0 {
				if update.Message.Text == "/get" {
					ph := photo.New()
					ph.Send(update.Message.Chat.Id)
				}
			}
		}
	}()
}
