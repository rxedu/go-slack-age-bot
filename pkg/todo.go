package pkg

import (
	"fmt"
	"github.com/rxedu/go-slack-age-bot/v1/internal"
)

func PrintMessage() {
	fmt.Println(internal.Message)
}

func GetMessage() string {
	return internal.Message
}
