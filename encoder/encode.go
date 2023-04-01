package encoder

import (
	"bytes"
	"encoding/json"
	"log"
)

func EncodeToJSONBuffer(obj any) *bytes.Buffer {
	json, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewBuffer(json)
	return body
}
