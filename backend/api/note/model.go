package note

import (
	"time"
	"github.com/google/uuid"
)

func NewNote(title string, content string, userID string, latitude float32, longitude float32) *Note {
	return &Note{
		ID:      uuid.New().String(),
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
