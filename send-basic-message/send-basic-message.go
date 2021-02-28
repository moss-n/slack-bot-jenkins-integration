package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"C01ACEBRWUC",
		slack.MsgOptionText("Hello world!", false),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}