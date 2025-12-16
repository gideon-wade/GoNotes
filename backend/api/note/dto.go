package note

import "time"

type NoteResponseDTO struct {
	ID      string    `json:"id" binding:"required"`
	Title   string    `json:"title" binding:"required"`
	Content string    `json:"content" binding:"required"`
	Date    time.Time `json:"date" binding:"required"`
}

type NewNoteRequestDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
