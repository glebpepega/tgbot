package photo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/glebpepega/goodvibesbot/decoder"
	"github.com/glebpepega/goodvibesbot/encoder"
	"github.com/glebpepega/goodvibesbot/link"
)

type Photo struct {
	Chat_id      int                  `json:"chat_id"`
	Photo        string               `json:"photo"`
	Caption      string               `json:"caption"`
	Reply_Markup InlineKeyboardMarkup `json:"reply_markup"`
}

type InlineKeyboardMarkup struct {
	Inline_Keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text          string `json:"text"`
	Callback_Data string `json:"callback_data"`
}

type Quote struct {
	Quote  string
	Author string
}

func New() *Photo {
	return &Photo{}
}

func (p *Photo) Send(chatID int) {
	p.Photo = getImgFromInternet()
	p.Chat_id = chatID
	p.Caption = getQuoteFromInternet()
	kb := p.newKeyboard()
	row := []InlineKeyboardButton{}
	button := InlineKeyboardButton{
		Text:          "🐶",
		Callback_Data: "🐶",
	}
	row = append(row, button)
	kb = append(kb, row)
	p.Reply_Markup.Inline_Keyboard = kb
	body := encoder.EncodeToJSONBuffer(p)
	_, err := http.Post(link.Link()+"/sendPhoto", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
}

func getQuoteFromInternet() string {
	quote := []Quote{}
	jsonResp, err := http.Get("https://api.breakingbadquotes.xyz/v1/quotes")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := jsonResp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	decoder.DecodeFromJSON(jsonResp.Body, &quote)
	result := fmt.Sprintf("\"%s\" — %s", quote[0].Quote, quote[0].Author)
	return result
}

func getImgFromInternet() string {
	var url []string
	jsonResp, err := http.Get("http://shibe.online/api/shibes?count=1&urls=true&httpsUrls=true")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := jsonResp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	decoder.DecodeFromJSON(jsonResp.Body, &url)
	return url[0]
}

func (p *Photo) newKeyboard() [][]InlineKeyboardButton {
	return [][]InlineKeyboardButton{}
}
