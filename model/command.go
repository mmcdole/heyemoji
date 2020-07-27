package model

import "github.com/shomali11/slacker"

// Command defines a command to be register to slack
type Command struct {
	Usage       string
	Description string
	Handler     func(ctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)
}
