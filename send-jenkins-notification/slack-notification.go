package main

import (
    "fmt"
    "os"
	"github.com/slack-go/slack"
)

func main() {
    api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
    args := os.Args[1:]
    fmt.Println(args)
	preText := "*Hello! Your Jenkins build has finished!*"
	jenkinsURL := "*Build URL:* " + args[0]
	buildResult  := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}
	dividerSection := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText + "\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)
	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, jenkinsBuildDetailsField)


	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)
	msg := slack.MsgOptionBlocks(
		preTextSection, 
		dividerSection,
		fieldsSection,
	)
	_, _, _, err := api.SendMessage(
		//channel or user ID that you want to send the message to
		"C01ACEBRWUC" ,
		msg,
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

}