package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxp-629997071139-630017779906-652868024880-2f136fcd4a9a644b662a583f62699c96")
	user, err := api.GetUserByEmail("tantatorn.s@kmi.tl")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
