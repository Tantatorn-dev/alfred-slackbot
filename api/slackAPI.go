package api

import (
	"github.com/nlopes/slack"
)

//SlackClient api
type SlackClient struct {
	API *slack.Client
}

//Init slack api initialisationS
func (client *SlackClient) Init() {
	TOKEN := "xoxb-629997071139-642100825234-KOXW3j2IWEFzAe0tBJ7AaLRi"
	client.API = slack.New(TOKEN)
}
