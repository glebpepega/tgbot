package fetcher

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/glebpepega/goodvibesbot/decoder"
	"github.com/glebpepega/goodvibesbot/encoder"
	"github.com/glebpepega/goodvibesbot/link"
	"github.com/glebpepega/goodvibesbot/update"
)

func GetUpdates(structResp *update.UpdateResponse, offset *int, ch chan update.Update) {
	for {
		func() {
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
		}()
	}
}
