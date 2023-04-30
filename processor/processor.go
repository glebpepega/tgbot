package processor

import (
	"github.com/glebpepega/goodvibesbot/sendable/photo"
	"github.com/glebpepega/goodvibesbot/sendable/text"
	"github.com/glebpepega/goodvibesbot/update"
)

func Process(u update.Update) {
	if u.Message.Text == "/get" {
		SendPic(u.Message.Chat.ID)
	} else {
		SendText(u.Message.Chat.ID, "Click /get")
	}
	if u.Callback_Query.ID != "" {
		u.Callback_Query.Answer()
		SendPic(u.Callback_Query.From.ID)
	}
}

func SendPic(id int) {
	ph := photo.New()
	ph.Send(id)
}

func SendText(id int, message string) {
	tx := text.New()
	tx.Send(id, message)
}
