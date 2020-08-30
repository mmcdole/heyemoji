package database

import "time"

func NewJSONLineDriver(path string) Driver {
	return &JSONLineDriver{path: path}
}

type JSONLineDriver struct {
	path string
}

func (d JSONLineDriver) Open() error {
	return nil

}

func (d JSONLineDriver) GiveKarma(to string, from string, amount int, date time.Time) error {
	return nil
}

func (d JSONLineDriver) QueryKarmaGiven(user string, since time.Time) (int, error) {
	return 0, nil

}

func (d JSONLineDriver) QueryKarmaReceived(user string, since time.Time) (int, error) {
	return 0, nil

}

func (d JSONLineDriver) QueryLeaderboard(since time.Time) (bool, error) {
	return false, nil
}
