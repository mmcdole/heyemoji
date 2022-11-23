module github.com/mmcdole/heyemoji

go 1.14

require (
	github.com/google/uuid v1.1.2
	github.com/shomali11/slacker v0.0.0-20200610181250-3156f073f291
	github.com/slack-go/slack v0.11.4
	github.com/spf13/viper v1.7.0
	golang.org/x/sys v0.2.0 // indirect
)

replace github.com/shomali11/slacker => github.com/mmcdole/slacker v0.0.0-20200726060059-aa5253e0331c
