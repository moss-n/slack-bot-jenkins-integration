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
	jenkinsURL := args[0]
	buildResult := args[1]
	buildNumber := args[2]
	jobName := args[3]
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult
	//send attachment
	attachment := slack.Attachment{
		Pretext: "Hello! Your Jenkins build has finished!",
		Text:    "Jenkins Build Details: ",
		
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: jenkinsBuildDetails,
					Value: "Build URL: " + jenkinsURL,
				},
			},
		}

	
	channelID, timestamp, err := api.PostMessage(
		//channel or user ID that you want to send the message to
		"C01ACEBRWUC" ,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

}