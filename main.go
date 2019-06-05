package main

import (
	"alfred-slackbot/api"
)

func main() {

	var client api.SlackClient

	client.Init()
	client.SendMessage("random", "test slack bot")

}
