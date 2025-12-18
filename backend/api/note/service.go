package note

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateNewNote(newNoteRequest NewNoteRequestDTO) (NoteResponseDTO, error) {
	newNote := NewNote(newNoteRequest.Title, newNoteRequest.Content, newNoteRequest.UserID, newNoteRequest.Latitude, newNoteRequest.Longitude)
	err := s.repo.Save(newNote)
	noteResponse := NoteResponseDTO{
		ID:      newNote.ID,
		Title:   newNote.Title,
		Content: newNote.Content,
		Date:    newNote.Date,
	}
	return noteResponse, err
}
