package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {

	TOKEN := "xoxb-629997071139-642100825234-KOXW3j2IWEFzAe0tBJ7AaLRi"

	api := slack.New(TOKEN)
	attachment := slack.Attachment{
		Pretext: "Test",
		Text:    "Test Test Test",
	}

	channelID, timestamp, err := api.PostMessage("random", slack.MsgOptionText("Some text", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
