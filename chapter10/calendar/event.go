package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

// Get the title of the event.
func (e *Event) Title() string {
	return e.title
}

// Set the title of the event. Must be 0 >= 1 <= 50 characters.
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) == 0 || utf8.RuneCountInString(title) > 50 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

//get title length
func (e *Event) TitleLength() int {
	return utf8.RuneCountInString(e.title)
}
