package api

import (
	"alfred-slackbot/tokens"
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

//SlackClient api
type SlackClient struct {
	API *slack.Client
}

//Init slack api initialisation
func (client *SlackClient) Init() {
	client.API = slack.New(
		tokens.BotUserOAuthAccessToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)))
}

//SendMessage a method for sending a message
func (client *SlackClient) SendMessage(channel string, message string) {
	channelID, timestamp, err := client.API.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}

//GetChannels get slack channels
func (client *SlackClient) GetChannels() []string {
	var channelsArr []string

	channels, err := client.API.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return []string{}
	}

	for _, channel := range channels {
		channelsArr = append(channelsArr, channel.Name)
	}

	return channelsArr
}
