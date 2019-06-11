package main

import (
	slack "alfred-slackbot/api/slackAPI"
	slash "alfred-slackbot/api/slashCommand"
	"alfred-slackbot/tokens"
)

func main() {
	tokens.LoadEnv()

	var client slack.SlackClient

	client.Init()
	slash.Serve()
}
