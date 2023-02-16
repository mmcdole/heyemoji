package event

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/mmcdole/heyemoji/database"
	"github.com/slack-go/slack"
)

const (
	maxLeaderEntries = 10
)

func NewLeaderHandler(db database.Driver) LeaderHandler {
	return LeaderHandler{db: db}
}

type LeaderHandler struct {
	db database.Driver
}

func (h LeaderHandler) Matches(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, ok := e.Data.(*slack.MessageEvent)
	if !ok {
		return false
	}
	if !IsBotMentioned(msg, rtm) && !IsDirectMessage(msg) {
		return false
	}
	if strings.Contains(strings.ToLower(msg.Text), "leaderboard") {
		return true
	}
	return false
}

func (h LeaderHandler) Execute(e slack.RTMEvent, rtm *slack.RTM) bool {
	ev, _ := e.Data.(*slack.MessageEvent)

	var header string
	start := time.Now()
	if strings.Contains(ev.Text, "day") {
		start = start.AddDate(0, 0, -1)
		header = "Today's Leaderboard"
	} else if strings.Contains(ev.Text, "week") {
		start = start.AddDate(0, 0, -7)
		header = "This Week's Leaderboard"
	} else if strings.Contains(ev.Text, "year") {
		start = start.AddDate(0, 0, -365)
		header = "This Year's Leaderboard"
	} else if strings.Contains(ev.Text, "quarter") {
		start = start.AddDate(0, 0, -91)
		header = "This Quarter's Leaderboard"
	} else if strings.Contains(ev.Text, "all") {
		start = start.AddDate(-99, 0, 0)
		header = "All Time Leaderboard"
	} else {
		/* Default to Month */
		start = start.AddDate(0, 0, -30)
		header = "This Month's Leaderboard"
	}

	leaders, err := h.db.QueryLeaderboard(start)
	if err != nil {
		return false
	}

	if len(leaders) == 0 {
		h.handleEmptyLeaderboard(ev, rtm)
		return true
	}

	h.handleSuccess(ev, rtm, leaders, header)
	return true
}

func (h LeaderHandler) handleSuccess(ev *slack.MessageEvent, rtm *slack.RTM, leaders map[string]int, header string) error {
	rank := h.rankMapStringInt(leaders)
	msg := fmt.Sprintf(">*%s*\n", header)
	for i := 0; i < len(rank) && i < maxLeaderEntries; i++ {
		name := rank[i]
		uinfo, err := rtm.GetUserInfo(rank[i])
		if err == nil {
			name = uinfo.RealName
		}
		msg += fmt.Sprintf(">%d) %s `%d`\n", i+1, name, leaders[rank[i]])
	}
	msg += ">\n"
	msg += "> You can view other leaderboards! :tada:\n"
	msg += "> *leaderboard <day | week | month>*"

	rtm.SendMessage(rtm.NewOutgoingMessage(msg, ev.Channel))
	return nil
}

func (h LeaderHandler) handleEmptyLeaderboard(ev *slack.MessageEvent, rtm *slack.RTM) error {
	_, err := rtm.PostEphemeral(
		ev.Channel,
		ev.User,
		slack.MsgOptionText("Nobody has given any emoji points yet!", false),
	)
	return err
}

func (h LeaderHandler) rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
