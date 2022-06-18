package main

import (
	"context"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-525672713012-3667426336609-TyK6bOLO5gTOf4TYQPGcuVRb")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03K8LJ2K9R-3640167824327-74ca9f70602799072ff67e4d12d0a5bb18839de3b3145154653478cd11cf78c3")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	definition := &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}

	bot.Command("ping", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
