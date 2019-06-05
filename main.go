package main

import (
	"alfred-slackbot/api"
	"fmt"

	"github.com/nlopes/slack"
)

func main() {

	var client api.SlackClient

	client.Init()

	attachment := slack.Attachment{
		Pretext: "Test",
		Text:    "Test Test Test",
	}

	channelID, timestamp, err := client.API.PostMessage("random", slack.MsgOptionText("Some text", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)

}
