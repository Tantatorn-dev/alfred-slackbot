package api

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

//Serve the slash commands server
func Serve() {
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

	if !s.ValidateToken("kjD6xcFNX77d93VOxYbDu4sq") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/show_channels":
		response := fmt.Sprintf("hello moto")
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
