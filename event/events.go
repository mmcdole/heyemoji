package event

import (
	"github.com/mmcdole/heyemoji/model"
	"github.com/shomali11/slacker"
)

var commands []model.Event

func Register(usage string, description string, handler func(ctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)) {
	commands = append(commands, model.Command{Usage: usage, Description: description, Handler: handler})
}

func EventHandler() []model.Event {
	return events
}
