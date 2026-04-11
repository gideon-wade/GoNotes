package note

import (
	"testing"
)

func TestCreateNewNote(t *testing.T) {
	noteRepo := NewInMemNoteRepository()
	noteService := NewService(noteRepo)

	_, err := noteService.CreateNewNote(NewNoteRequestDTO{
		Title:     "Test Note",
		Content:   "This is a test note.",
		Latitude:  37.7749,
		Longitude: -122.4194,
	}, "user123")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
