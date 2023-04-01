package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/glebpepega/goodvibesbot/decoder"
	"github.com/glebpepega/goodvibesbot/update"
	"github.com/joho/godotenv"
)

var (
	url string
)

func main() {
	err := godotenv.Load("X:/mycode/goodvibesbot/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	tgToken := os.Getenv("t")
	url = "https://api.telegram.org/bot" + tgToken
	structResp := update.NewResponse()
	offset := 0
	updateChan := make(chan update.Update)

	go func() {
		for v := range updateChan {
			fmt.Println(v)
		}
	}()

	for {
		getUpdates(structResp, &offset, updateChan)
	}

}

func getUpdates(structResp *update.UpdateResponse, offset *int, ch chan update.Update) {
	offsetStr := fmt.Sprintf(`{"offset":"%s"}`, strconv.Itoa(*offset))
	offsetJSON, err := json.Marshal(offsetStr)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewBuffer(offsetJSON)
	jsonResp, err := http.Post(url+"/getUpdates", "application/json", body)
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
