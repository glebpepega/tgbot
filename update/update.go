package update

import (
	"log"
	"net/http"

	"github.com/glebpepega/goodvibesbot/encoder"
	"github.com/glebpepega/goodvibesbot/link"
)

type Update struct {
	Update_id      int
	Message        Message
	Callback_Query CallbackQuery
}

type Message struct {
	Chat Chat
	Text string
}

type Chat struct {
	ID int
}

type CallbackQuery struct {
	ID   string
	From User
}

type User struct {
	ID         int
	Is_Bot     bool
	First_Name string
}

type CallbackQueryAnswer struct {
	Callback_query_id string `json:"callback_query_id"`
	Text              string `json:"text"`
}

func NewUpdate() *Update {
	return &Update{}
}

func (cq *CallbackQuery) Answer() {
	cbqa := CallbackQueryAnswer{
		Callback_query_id: cq.ID,
		Text:              "🥰",
	}
	body := encoder.EncodeToJSONBuffer(cbqa)
	resp, err := http.Post(link.Link()+"/answerCallbackQuery", "application/json", body)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
}
