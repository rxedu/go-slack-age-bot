package main

import (
	"os"

	"github.com/rxedu/go-slack-age-bot"
)

func main() {
	goslackagebot.StartBot(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
}
