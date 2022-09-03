package main

import (
	"os"

	"github.com/rxedu/go-slack-age-bot/v1/pkg"
)

func main() {
	pkg.StartBot(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
}
