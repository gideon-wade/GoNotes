package note

import (
	"time"
)

func NewNote(title string, description string) *Note {
	return &Note{
		ID:          "d14a0232-8f22-44ce-b8c5-1d51053a361d",
		Title:       title,
		Description: description,
		Date:        time.Now(),
	}
}

type Note struct {
	ID          string
	Title       string
	Description string
	Date        time.Time
}
