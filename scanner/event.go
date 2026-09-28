package scanner

import "time"

type Event struct {
	Description string
	Userinput   string
	DateAt      time.Time
}

func NewEvent(description string, userInput string) Event {
	return Event{
		Description: description,
		Userinput:   userInput,
		DateAt:      time.Now(),
	}
}
