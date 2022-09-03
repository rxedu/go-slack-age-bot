package goslackagebot

import (
	"context"
	"fmt"
	"log"

	"github.com/shomali11/slacker"

	"github.com/rxedu/go-slack-age-bot/internal"
)

func StartBot(botToken string, appToken string) {
	bot := slacker.NewClient(botToken, appToken)
	go printCommandEvents(bot.CommandEvents())

	internal.RegisterCommands(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
