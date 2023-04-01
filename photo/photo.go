package photo

import (
	"log"
	"net/http"

	"github.com/glebpepega/goodvibesbot/decoder"
	"github.com/glebpepega/goodvibesbot/encoder"
	"github.com/glebpepega/goodvibesbot/link"
)

type Photo struct {
	Chat_id int    `json:"chat_id"`
	Photo   string `json:"photo"`
}

func New() *Photo {
	return &Photo{}
}

func (p *Photo) Send(chatID int) {
	p.Photo = getFromInternet()
	p.Chat_id = chatID
	body := encoder.EncodeToJSONBuffer(p)
	_, err := http.Post(link.Link()+"/sendPhoto", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
}

func getFromInternet() string {
	var url []string
	jsonResp, err := http.Get("http://shibe.online/api/shibes?count=1&urls=true&httpsUrls=true")
	if err != nil {
		log.Fatal(err)
	}
	decoder.DecodeFromJSON(jsonResp.Body, &url)
	return url[0]
}
