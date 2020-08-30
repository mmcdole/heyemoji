package database

import "time"

type Driver interface {
	Open() error
	GiveKarma(to string, from string, amount int, emoji string, date time.Time) error
	QueryKarmaGiven(user string, since time.Time) (int, error)
	QueryKarmaReceived(user string, since time.Time) (int, error)
	QueryLeaderboard(since time.Time) (bool, error)
}
