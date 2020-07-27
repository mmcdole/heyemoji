package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	BotName       string
	DatabasePath  string
	SlackToken    string
	SlackEmoji    string
	SlackDailyCap int
	WebSocketPort int
}

func readConfig() *Config {
	viper.SetEnvPrefix("hey")
	viper.AutomaticEnv()

	viper.SetDefault("bot_name", "heyemoji")
	viper.SetDefault("database_path", "./data/")
	viper.SetDefault("slack_api_token", "")
	viper.SetDefault("slack_emoji", ":star:")
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

	return c
}
