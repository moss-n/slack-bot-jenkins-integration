package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
	"net/http"
)

func sendSlackMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Sent Slack Message!</h1>")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	preText := "*Hello! Your Jenkins build has finished!*"
	dividerSection1 := slack.NewDividerBlock() 
	preTextField := slack.NewTextBlockObject("mrkdwn", preText + "\n\n", false, false) 
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
			preTextSection,
			dividerSection1,
	)

	_, _, _, err := api.SendMessage(
		"C01ACEBRWUC",
		msg,
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
  }

func main() {

	http.HandleFunc("/sendSlackMessage", sendSlackMessage)
	http.ListenAndServe(":8091", nil)

	
}