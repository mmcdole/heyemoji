package main

import (
	"context"
	"log"

	"github.com/mmcdole/heyemoji/cmd"

	"github.com/shomali11/slacker"
)

func main() {

	cfg := readConfig()

	bot := slacker.NewClient(cfg.SlackToken)

	for _, command := range cmd.CommandList() {
		bot.Command(command.Usage, &slacker.CommandDefinition{
			Description: command.Description,
			Handler:     command.Handler,
		})
	}

	event.EventList()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
