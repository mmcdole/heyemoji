package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/heyemoji/database"
	"github.com/mmcdole/heyemoji/event"
	"github.com/slack-go/slack"
)

func main() {

	cfg := readConfig()
	db := database.NewJSONLineDriver(cfg.DatabasePath, 5000)

	if err := db.Open(); err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}

	handlers := []event.EventHandler{
		event.PingHandler{},
		event.NewLeaderHandler(db),
		event.NewHelpHandler(cfg.SlackDailyCap, cfg.SlackEmojiMap),
		event.NewPointsHandler(cfg.SlackDailyCap, db),
		event.NewEmojiHandler(cfg.SlackEmojiMap, cfg.SlackDailyCap, db),
	}

	api := slack.New(
		cfg.SlackToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		fmt.Printf("Message: %v\n", msg.Data)
		if event.IsBotMessage(msg) {
			continue
		}
		for _, h := range handlers {
			if !h.Matches(msg, rtm) {
				continue
			}
			if handled := h.Execute(msg, rtm); handled {
				break
			}
		}
	}
}
