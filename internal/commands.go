package internal

import (
	"fmt"
	"strconv"

	"github.com/shomali11/slacker"
)

func RegisterCommands(bot *slacker.Slacker) {
	bot.Command("I was born in <year>", &slacker.CommandDefinition{
		Description: "About how old are you?",
		Examples:    []string{"I was born in <year>"},
		Handler:     yobHandler,
	})
}

func yobHandler(ctx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	year := r.Param("year")
	yob, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
		return
	}
	age := 2021 - yob
	msg := fmt.Sprintf("You are about %d years old.", age)
	err = w.Reply(msg)
	if err != nil {
		fmt.Println(err)
	}
}
