package event

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/heyemoji/database"

	"github.com/slack-go/slack"
)

var userRegex = regexp.MustCompile("<@[A-Z0-9]{2,}>")

func NewEmojiHandler(emoji map[string]int, dailyCap int, db database.Driver) EmojiHandler {
	return EmojiHandler{emoji: emoji, dailyCap: dailyCap, db: db}
}

type EmojiHandler struct {
	emoji    map[string]int
	db       database.Driver
	dailyCap int
	re       *regexp.Regexp
}

func (h EmojiHandler) Matches(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, ok := e.Data.(*slack.MessageEvent)
	if !ok {
		return false
	}
	// Message contains one of the pre-defined slack emojis
	if re := h.emojiRegex(); re.MatchString(msg.Text) {
		return true
	}
	return false
}

func (h EmojiHandler) Execute(e slack.RTMEvent, rtm *slack.RTM) bool {
	ev, _ := e.Data.(*slack.MessageEvent)

	emojis := h.parseEmojis(ev.Text)
	if len(emojis) <= 0 {
		return false
	}

	users := h.parseUsers(ev.Text)
	if len(users) <= 0 {
		h.handleNoRecipient(ev, rtm, emojis)
		return true
	}

	if IsDirectMessage(ev) {
		h.handleDirectMessage(ev, rtm)
		return true
	}

	required := h.getRequiredKarma(emojis, len(users))
	balance := h.getKarmaBalance(ev.User)
	if balance < required {
		h.handleInsufficientKarma(ev, rtm, balance, required, emojis, users)
		return true
	}

	filteredUsers := Filter(users, func(val string) bool {
		return val != ev.User
	})
	if len(filteredUsers) != len(users) {
		h.handleSelfKarma(ev, rtm)
	}

	if len(filteredUsers) > 0 {
		h.handleSuccess(ev, rtm, emojis, filteredUsers)
		return true
	}

	return false
}

func (h EmojiHandler) handleSuccess(ev *slack.MessageEvent, rtm *slack.RTM, emojis []string, users []string) error {

	karmaPerUser := h.getRequiredKarma(emojis, 1)

	for _, user := range users {
		h.db.GiveKarma(user, ev.User, karmaPerUser, time.Now())
	}

	balance := h.getKarmaBalance(ev.User)
	msg := fmt.Sprintf("%s received *%d emoji point(s)* from you.  You have *%d point(s)* left to give out today",
		strings.Join(users, " "),
		karmaPerUser,
		balance)

	_, _, err := rtm.PostMessage(
		ev.User,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(true),
	)
	return err
}

func (h EmojiHandler) handleSelfKarma(ev *slack.MessageEvent, rtm *slack.RTM) error {
	_, err := rtm.PostEphemeral(
		ev.Channel,
		ev.User,
		slack.MsgOptionText("Sorry, you can only give emoji points to other people on your team.", false),
	)
	return err
}

func (h EmojiHandler) handleDirectMessage(ev *slack.MessageEvent, rtm *slack.RTM) error {
	_, _, err := rtm.PostMessage(
		ev.Channel,
		slack.MsgOptionText("Sorry, you can only give people emoji points in channels.", false),
	)
	return err
}

func (h EmojiHandler) handleNoRecipient(ev *slack.MessageEvent, rtm *slack.RTM, emojis []string) error {
	msg := fmt.Sprintf("Give someone that :%s: by adding it after their username, like this: @username :%s:", emojis[0], emojis[0])

	_, err := rtm.PostEphemeral(
		ev.Channel,
		ev.User,
		slack.MsgOptionText(msg, false),
	)
	return err
}

func (h EmojiHandler) handleInsufficientKarma(ev *slack.MessageEvent, rtm *slack.RTM, balance int, required int, emojis []string, users []string) error {
	msg := fmt.Sprintf("Whoops! You tried to give *%d* emoji point(s). "+
		"You have *%d* point(s) left to give today. "+
		"Your point balance will reset in *%s*.",
		required,
		balance,
		h.fmtDuration(h.timeTillReset()))

	_, err := rtm.PostEphemeral(
		ev.Channel,
		ev.User,
		slack.MsgOptionText(msg, false),
	)
	return err
}

func (h EmojiHandler) emojiRegex() *regexp.Regexp {
	if h.re != nil {
		return h.re
	}

	emojis := make([]string, 0, len(h.emoji))
	for k := range h.emoji {
		emojis = append(emojis, k)
	}

	// Add colons to emoji
	for i, e := range emojis {
		emojis[i] = fmt.Sprintf(":%s:", e)
	}

	// Build OR regex of emoji
	str := strings.Join(emojis, "|")
	h.re = regexp.MustCompile(str)
	return h.re
}

func (h EmojiHandler) parseUsers(msg string) []string {
	users := userRegex.FindAllString(msg, -1)
	for i, u := range users {
		u = strings.Replace(u, "<@", "", -1)
		u = strings.Replace(u, ">", "", -1)
		users[i] = u
	}
	return users
}

func (h EmojiHandler) parseEmojis(msg string) []string {
	re := h.emojiRegex()
	emoji := re.FindAllString(msg, -1)
	for i, e := range emoji {
		emoji[i] = strings.Replace(e, ":", "", -1)
	}
	return emoji
}

// Get the daily karma balance left for a user
func (h EmojiHandler) getKarmaBalance(user string) int {
	given, _ := h.db.QueryKarmaGiven(user, h.lastReset())
	return h.dailyCap - given
}

// Get the required karma for a set of emoji and receivers
func (h EmojiHandler) getRequiredKarma(emoji []string, numUsers int) int {
	total := 0
	for _, e := range emoji {
		if karma, ok := h.emoji[e]; ok {
			total += karma
		}
	}
	return total * numUsers
}

// Format a time.Duration till karma reset for display to user
func (h EmojiHandler) fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	hr := d / time.Hour
	d -= hr * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%2d hours and %2d minutes", hr, m)
}

// Return last emoji reset time
func (h EmojiHandler) lastReset() time.Time {
	reset := time.Now()
	reset = time.Date(reset.Year(), reset.Month(), reset.Day(), 0, 0, 0, 0, reset.Location())
	return reset
}

// Return next emoji reset time
func (h EmojiHandler) nextReset() time.Time {
	now := time.Now()
	reset := now.AddDate(0, 0, 1)
	reset = time.Date(reset.Year(), reset.Month(), reset.Day(), 0, 0, 0, 0, reset.Location())
	return reset
}

// Return time till the next emoji reset
func (h EmojiHandler) timeTillReset() time.Duration {
	reset := h.nextReset()
	// Duration till reset
	return reset.Sub(time.Now())
}
