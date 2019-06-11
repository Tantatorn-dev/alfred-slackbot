package tokens

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//BotUserOAuthAccessToken OAuth token for a bot user
var BotUserOAuthAccessToken string

//VerificationToken a token for verification
var VerificationToken string

//LoadEnv load environment.env
func LoadEnv() {
	err := godotenv.Load("environment.env")
	if err != nil {
		log.Fatal("error load .env file")
	}

	VerificationToken = os.Getenv("SLACK_VERIFICATION_TOKEN")
	BotUserOAuthAccessToken = os.Getenv("SLACK_BOTUSEROAUTH_TOKEN")
}
