package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type message struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

func makeMessage(text string) message {
	channel := os.Getenv("SLACK_CHANNEL")
	username := os.Getenv("SLACK_USERNAME")
	iconEmoji := os.Getenv("SLACK_ICONEMOJI")

	return message{
		Channel:   channel,
		Username:  username,
		Text:      text,
		IconEmoji: iconEmoji,
	}
}

func Publish(text string) string {
	url := os.Getenv("SLACK_URL")

	if url == "" {
		panic("SLACK_URL needs to be specified")
	}

	data, _ := json.Marshal(makeMessage(text))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
