package slash

import (
	api "alfred-slackbot/api/slackAPI"
	"alfred-slackbot/tokens"
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

var apiPtr *api.SlackClient

//Serve the slash commands server
func Serve(ptr *api.SlackClient) {
	apiPtr = ptr

	http.HandleFunc("/receive", slashCommandHandler)

	fmt.Println("Server Listening")
	http.ListenAndServe(":8080", nil)
}

func slashCommandHandler(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(tokens.VerificationToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/show_channels":
		response := fmt.Sprintln(apiPtr.GetChannels())
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
