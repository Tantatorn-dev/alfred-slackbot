package main

import (
	"alfred-slackbot/api"
	"fmt"
)

func main() {

	var client api.SlackClient

	client.Init()
	client.SendMessage("random", "test slack bot")
	fmt.Println(client.GetChannels())
}
