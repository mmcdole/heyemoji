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
	db := database.NewJSONLineDriver(cfg.DatabasePath)

	if err := db.Open(); err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}

	events := []event.EventHandler{
		event.PingHandler{},
		event.HelpHandler{},
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

		for _, ev := range events {
			if !ev.Matches(msg, rtm) {
				continue
			}
			if handled := ev.Execute(msg, rtm); handled {
				break
			}
		}
	}
}
