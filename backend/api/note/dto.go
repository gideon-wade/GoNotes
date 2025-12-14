package note

import "time"


type NoteResponseDTO struct {
	ID          string    `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

type NewNoteRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}



