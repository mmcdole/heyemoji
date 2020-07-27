package cmd

import "github.com/shomali11/slacker"

func init() {
	Register("ping", "Ping test command!", ping)
}

func ping(ctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	response.Reply("pong!")
}
