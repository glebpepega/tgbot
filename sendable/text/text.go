package text

import (
	"log"
	"net/http"

	"github.com/glebpepega/goodvibesbot/encoder"
	"github.com/glebpepega/goodvibesbot/link"
)

type Text struct {
	Chat_id int    `json:"chat_id"`
	Text    string `json:"text"`
}

func New() *Text {
	return &Text{}
}

func (t *Text) Send(chatID int, message string) {
	t.Text = message
	t.Chat_id = chatID
	body := encoder.EncodeToJSONBuffer(t)
	resp, err := http.Post(link.Link()+"/sendMessage", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	if err := resp.Body.Close(); err != nil {
		log.Fatal(err)
	}
}
