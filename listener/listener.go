package listener

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/glebpepega/goodvibesbot/decoder"
	"github.com/glebpepega/goodvibesbot/link"
	"github.com/glebpepega/goodvibesbot/processor"
	"github.com/glebpepega/goodvibesbot/update"
)

func ListenAndReactToUpdates(webHookURL string) {
	SetWebHook(webHookURL)
	http.HandleFunc("/", extractAndProcessUpdates)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func extractAndProcessUpdates(w http.ResponseWriter, r *http.Request) {
	u := update.NewUpdate()
	decoder.DecodeFromJSON(r.Body, u)
	processor.Process(*u)
}

func SetWebHook(webHookURL string) {
	body := bytes.NewBuffer([]byte(fmt.Sprintf(`{"url":"%s"}`, webHookURL)))
	resp, err := http.Post(link.Link()+"/setWebhook", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	if err := resp.Body.Close(); err != nil {
		log.Println(err)
	}
}
