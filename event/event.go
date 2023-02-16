package event

import (
	"fmt"
	"strings"
	"time"

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

func IsBotMessage(msg slack.RTMEvent) bool {
	msgEvent, ok := msg.Data.(*slack.MessageEvent)
	if !ok {
		return true
	}
	return msgEvent.BotID != ""
}

// Get last point reset time
func LastPointReset() time.Time {
	reset := time.Now()
	reset = time.Date(reset.Year(), reset.Month(), reset.Day(), 0, 0, 0, 0, reset.Location())
	return reset
}

// Get next point reset time
func NextPointReset() time.Time {
	now := time.Now()
	reset := now.AddDate(0, 0, 1)
	reset = time.Date(reset.Year(), reset.Month(), reset.Day(), 0, 0, 0, 0, reset.Location())
	return reset
}

// Get time till the next emoji reset
func TimeTillPointReset() time.Duration {
	reset := NextPointReset()
	// Duration till reset
	return reset.Sub(time.Now())
}

// Format a time.Duration till karma reset for display to user
func FmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	hr := d / time.Hour
	d -= hr * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%2d hours and %2d minutes", hr, m)
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

func Keys(m map[string]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
