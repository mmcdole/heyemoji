package event

import (
	"fmt"
	"strings"

	"github.com/mmcdole/heyemoji/database"
	"github.com/slack-go/slack"
)

type PointsHandler struct {
	db       database.Driver
	dailyCap int
}

func NewPointsHandler(dailyCap int, db database.Driver) PointsHandler {
	return PointsHandler{db: db, dailyCap: dailyCap}
}

func (h PointsHandler) Matches(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, ok := e.Data.(*slack.MessageEvent)
	if !ok {
		return false
	}
	if !IsBotMentioned(msg, rtm) && !IsDirectMessage(msg) {
		return false
	}
	if strings.EqualFold(msg.Text, "points") {
		return true
	}
	return false
}

func (h PointsHandler) Execute(e slack.RTMEvent, rtm *slack.RTM) bool {
	ev, _ := e.Data.(*slack.MessageEvent)

	given, _ := h.db.QueryKarmaGiven(ev.User, LastPointReset())
	balance := h.dailyCap - given

	timeTillReset := FmtDuration(TimeTillPointReset())
	msg := fmt.Sprintf("You have %d emoji points left to give today. Your points will reset in %s.", balance, timeTillReset)
	rtm.SendMessage(rtm.NewOutgoingMessage(msg, ev.Channel))

	return true
}
