package model

import "github.com/shomali11/slacker"

// Command defines a command to be register to slack
type Event struct {
	Handler func(ctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)
}
