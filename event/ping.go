package event

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

type PingHandler struct{}

func (p PingHandler) Matches(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, ok := e.Data.(*slack.MessageEvent)
	if !ok {
		return false
	}
	if !IsBotMentioned(msg, rtm) && !IsDirectMessage(msg) {
		return false
	}
	if strings.Contains(strings.ToLower(msg.Text), "ping") {
		return true
	}
	return false
}

func (p PingHandler) Execute(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, _ := e.Data.(*slack.MessageEvent)

	fmt.Println("EXECUTE PING START")
	fmt.Printf("Channel: %s\n", msg.Channel)
	rtm.SendMessage(rtm.NewOutgoingMessage("pong", msg.Channel))

	return true
}
