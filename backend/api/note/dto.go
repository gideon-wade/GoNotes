package note

import "time"

type NoteResponseDTO struct {
	ID      string    `json:"id" binding:"required"`
	UserID  string    `json:"userID" binding:"required"`
	Title   string    `json:"title" binding:"required"`
	Content string    `json:"content" binding:"required"`
	Date    time.Time `json:"date" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}

type NewNoteRequestDTO struct {
	UserID  string `json:"userID" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}
