package event

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

type LeaderHandler struct{}

func (h LeaderHandler) Matches(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, ok := e.Data.(*slack.MessageEvent)
	if !ok {
		return false
	}
	if !IsBotMentioned(msg, rtm) && !IsDirectMessage(msg) {
		return false
	}
	if strings.Contains(msg.Text, "leaderboard") {
		return true
	}
	return false
}

func (h LeaderHandler) Execute(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, _ := e.Data.(*slack.MessageEvent)

	fmt.Println("EXECUTE PING START")
	fmt.Printf("Channel: %s\n", msg.Channel)
	rtm.SendMessage(rtm.NewOutgoingMessage("pong", msg.Channel))

	return true
}
