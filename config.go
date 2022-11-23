package main

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	BotName       string
	DatabasePath  string
	SlackToken    string
	SlackEmoji    string
	SlackEmojiMap map[string]int
	SlackDailyCap int
	WebSocketPort int
}

func readConfig() *Config {
	viper.SetEnvPrefix("hey")
	viper.AutomaticEnv()

	viper.SetDefault("bot_name", "heyemoji")
	viper.SetDefault("database_path", "./data/")
	viper.SetDefault("slack_token", "")
	viper.SetDefault("slack_emoji", "star:1")
	viper.SetDefault("slack_daily_cap", 5)
	viper.SetDefault("websocket_port", 3334)

	c := &Config{
		BotName:       viper.GetString("bot_name"),
		DatabasePath:  viper.GetString("database_path"),
		SlackToken:    viper.GetString("slack_token"),
		SlackEmoji:    viper.GetString("slack_emoji"),
		SlackDailyCap: viper.GetInt("slack_daily_cap"),
		WebSocketPort: viper.GetInt("websocket_port"),
	}

	c.SlackEmojiMap = createEmojiValueMap(c.SlackEmoji)

	return c
}

func createEmojiValueMap(e string) map[string]int {
	pairs := strings.Split(e, ",")

	evalues := make(map[string]int)
	for _, pair := range pairs {
		evalue := strings.Split(pair, ":")

		if len(evalue) != 2 {
			continue
		}

		emoji := evalue[0]
		if karma, err := strconv.Atoi(evalue[1]); err == nil {
			evalues[emoji] = karma
		}
	}
	return evalues
}
