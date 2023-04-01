package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/glebpepega/goodvibesbot/decoder"
	"github.com/glebpepega/goodvibesbot/encoder"
	"github.com/glebpepega/goodvibesbot/link"
	"github.com/glebpepega/goodvibesbot/photo"
	"github.com/glebpepega/goodvibesbot/update"
)

func main() {
	structResp := update.NewResponse()
	offset := 0
	updateChan := make(chan update.Update)

	go func() {
		for update := range updateChan {
			if len(update.Message.Entities) > 0 {
				if update.Message.Text == "/get" {
					ph := photo.New()
					ph.Send(update.Message.Chat.Id)
				}
			}
		}
	}()

	for {
		getUpdates(structResp, &offset, updateChan)
	}

}

func getUpdates(structResp *update.UpdateResponse, offset *int, ch chan update.Update) {
	offsetStr := fmt.Sprintf(`{"offset":"%s"}`, strconv.Itoa(*offset))
	body := encoder.EncodeToJSONBuffer(offsetStr)
	jsonResp, err := http.Post(link.Link()+"/getUpdates", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := jsonResp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	decoder.DecodeFromJSON(jsonResp.Body, structResp)
	if len(structResp.Result) != 0 && *offset < structResp.Result[len(structResp.Result)-1].Update_id+1 {
		if *offset != 0 {
			ch <- structResp.Result[len(structResp.Result)-1]
		}
		*offset = structResp.Result[len(structResp.Result)-1].Update_id + 1
	}
	time.Sleep(time.Millisecond * 500)
}
