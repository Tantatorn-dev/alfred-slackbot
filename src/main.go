package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("h0ZNuYmbt8d8UhOrUEiOBpzx")
	user, err := api.GetUserByEmail("tantatorn.s@kmi.tl")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
