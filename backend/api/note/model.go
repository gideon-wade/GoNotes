package note

import (
	"time"
)

func NewNote(title string, content string) *Note {
	return &Note{
		ID:      "d14a0232-8f22-44ce-b8c5-1d51053a361d",
		Title:   title,
		Content: content,
		Date:    time.Now(),
	}
}

type Note struct {
	ID      string
	Title   string
	Content string
	Date    time.Time
}
