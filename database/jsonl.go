package database

import (
	"bufio"
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
)

type KarmaEvent struct {
	ID     string    `json:"id"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
}

func NewJSONLineDriver(path string, maxDays int) Driver {
	return &JSONLineDriver{
		path:    path,
		maxDays: maxDays,
	}
}

type JSONLineDriver struct {
	maxDays int
	path    string
	cache   []KarmaEvent
	file    *os.File
}

func (d JSONLineDriver) Open() error {

	// Load karma events until 'maxDays' ago into memory
	maxDate := time.Now().AddDate(0, 0, -(d.maxDays + 1))
	events, err := d.loadEvents(d.path, maxDate)
	if err != nil {
		return err
	}
	d.cache = events

	// Open / create db file writing updates
	file, err := os.OpenFile(d.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	d.file = file

	return nil
}

func (d JSONLineDriver) Close() error {
	if d.file == nil {
		return nil
	}
	return d.file.Close()
}

func (d JSONLineDriver) GiveKarma(to string, from string, amount int, date time.Time) error {
	id := uuid.New()
	karma := KarmaEvent{ID: id.String(), From: from, To: to, Amount: amount, Date: date}

	json, err := json.Marshal(karma)
	if err != nil {
		return err
	}

	_, err = d.file.WriteString(string(json) + "\n")
	if err != nil {
		return err
	}

	d.cache = append(d.cache, karma)

	return nil
}

func (d JSONLineDriver) QueryKarmaGiven(user string, since time.Time) (int, error) {
	karma := 0
	for _, e := range d.getEvents() {
		if e.From == user && e.Date.After(since) {
			karma += e.Amount
		}
	}
	return karma, nil
}

func (d JSONLineDriver) QueryKarmaReceived(user string, since time.Time) (int, error) {
	karma := 0
	for _, e := range d.getEvents() {
		if e.To == user && e.Date.After(since) {
			karma += e.Amount
		}
	}
	return karma, nil
}

func (d JSONLineDriver) QueryLeaderboard(since time.Time) (map[string]int, error) {
	result := map[string]int{}
	for _, e := range d.getEvents() {
		if e.Date.After(since) {
			if val, ok := result[e.To]; ok {
				result[e.To] = val + e.Amount
			} else {
				result[e.To] = val
			}
		}
	}
	return result, nil
}

func (d JSONLineDriver) getEvents() []KarmaEvent {
	maxDate := time.Now().AddDate(0, 0, -(d.maxDays + 1))
	// Trim old events each time we access the field
	d.cache = d.removeEventsBefore(d.cache, maxDate)
	return d.cache
}

// Remove any days that are before a cutoff 'earliest' date
func (d JSONLineDriver) removeEventsBefore(events []KarmaEvent, earliest time.Time) []KarmaEvent {
	result := []KarmaEvent{}
	for _, event := range events {
		if event.Date.After(earliest) {
			result = append(result, event)
		}
	}
	return result
}

// Load events in the jsonl file that aren't before 'earliest' date
func (d JSONLineDriver) loadEvents(path string, earliest time.Time) ([]KarmaEvent, error) {

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	events := []KarmaEvent{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var event KarmaEvent
		json.Unmarshal(scanner.Bytes(), &event)
		if event.Date.After(earliest) {
			events = append(events, event)
		}
	}
	return events, scanner.Err()
}
