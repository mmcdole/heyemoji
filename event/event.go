package event

import (
	"github.com/slack-go/slack"
)

type EventHandler interface {
	Matches(slack.RTMEvent, *slack.RTM) bool
	Execute(slack.RTMEvent, *slack.RTM) bool
}

const (
	directChannelMarker = "D"
	userMentionFormat   = "<@%s>"
)
