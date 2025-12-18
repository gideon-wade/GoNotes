package note

import (
	"time"
)

func  NewNote(title string, content string, userID string, latitude float32, longitude float32) *Note {
	return &Note{
		ID:      "d14a0232-8f22-44ce-b8c5-1d51053a361d",
		UserID:  userID,
		Title:   title,
		Content: content,
		Date:    time.Now(),
		Latitude: latitude,
		Longitude: longitude,
	}
}

type Note struct {
	ID      	string
	UserID		string
	Title   	string
	Content 	string
	Date   		time.Time
	Latitude	float32
	Longitude	float32
}
