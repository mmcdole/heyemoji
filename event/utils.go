package event

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

func IsBotMentioned(event *slack.MessageEvent, rtm *slack.RTM) bool {
	info := rtm.GetInfo()
	return strings.Contains(event.Text, fmt.Sprintf(userMentionFormat, info.User.ID))
}

func IsDirectMessage(event *slack.MessageEvent) bool {
	return strings.HasPrefix(event.Channel, directChannelMarker)
}

func Filter(arr []string, cond func(string) bool) []string {
	result := []string{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
