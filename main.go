package main

import (
	"alfred-slackbot/api"
	"alfred-slackbot/tokens"
)

func main() {
	tokens.LoadEnv()

	var client api.SlackClient

	client.Init()

	client.SendMessage("random", "Hello")
}
