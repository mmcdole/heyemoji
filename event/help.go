package event

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/slack-go/slack"
)

type HelpHandler struct {
	dailyCap int
	emojiMap map[string]int
}

func NewHelpHandler(dailyCap int, emojiMap map[string]int) HelpHandler {
	return HelpHandler{dailyCap: dailyCap, emojiMap: emojiMap}
}

func (h HelpHandler) Matches(e slack.RTMEvent, rtm *slack.RTM) bool {
	msg, ok := e.Data.(*slack.MessageEvent)
	if !ok {
		return false
	}
	if !IsBotMentioned(msg, rtm) && !IsDirectMessage(msg) {
		return false
	}
	if strings.Contains(strings.ToLower(msg.Text), "help") {
		return true
	}
	return false
}

func (h HelpHandler) Execute(e slack.RTMEvent, rtm *slack.RTM) bool {
	tmp := `>*Directions*
>Add a recognition emoji after someone's username like this: *@username Great job!* :{{index .Emoji 0}}:. Everyone has {{.DailyCap}} emoji points to give out per day and can only give them in the channels I've been invited to.
>*Recognition Emoji*
>{{ range $key, $value := .EmojiMap }}:{{$key}}: *({{$value}} pts)*  {{end}}
>*Channel Commands*
>/invite <@{{.Botname}}>: to invite me to channels
><@{{.Botname}}> leaderboard <day|week|month>: to see the top 10 people on your leaderboard
><@{{.Botname}}> points: see how many emoji points you have left to give 
><@{{.Botname}}> help: get help with how to send recognition emoji 
>*Direct Message Commands*
>leaderboard <day|week|month>: to see the top 10 people on your leaderboard
>points: see how many emoji points you have left to give 
>help: get help with how to send recognition emoji`

	t := template.Must(template.New("help").Parse(tmp))

	var helpStr bytes.Buffer
	t.Execute(&helpStr, struct {
		Botname  string
		Emoji    []string
		EmojiMap map[string]int
		DailyCap int
	}{
		rtm.GetInfo().User.Name,
		Keys(h.emojiMap),
		h.emojiMap,
		h.dailyCap,
	})

	msg, _ := e.Data.(*slack.MessageEvent)
	rtm.SendMessage(rtm.NewOutgoingMessage(helpStr.String(), msg.Channel))

	return true
}
