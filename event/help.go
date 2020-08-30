package event

import (
	"github.com/slack-go/slack"
)

type HelpHandler struct{}

func (h HelpHandler) Matches(e slack.RTMEvent, rtm *slack.RTM) bool {
	return false
}

func (h HelpHandler) Execute(e slack.RTMEvent, rtm *slack.RTM) bool {
	return false
}
