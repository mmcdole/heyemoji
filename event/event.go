package event

import (
	"fmt"
	"strings"

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

func IsBotMentioned(event *slack.MessageEvent, rtm *slack.RTM) bool {
	info := rtm.GetInfo()
	return strings.Contains(event.Text, fmt.Sprintf(userMentionFormat, info.User.ID))
}

func IsDirectMessage(event *slack.MessageEvent) bool {
	return strings.HasPrefix(event.Channel, directChannelMarker)
}
