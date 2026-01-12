package note

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateNewNote(newNoteRequest NewNoteRequestDTO) (NoteResponseDTO, error) {
	newNote := NewNote(
		newNoteRequest.Title,
		newNoteRequest.Content,
		newNoteRequest.UserID,
		newNoteRequest.Latitude,
		newNoteRequest.Longitude,
	)
	err := s.repo.Save(*newNote)
	noteResponse := *newNote.ToNoteResponseDTO()
	return noteResponse, err
}

func (s *Service) GetAllNotes() ([]NoteResponseDTO, error) {
	notes, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	noteResponses := make([]NoteResponseDTO, len(notes))
	for i, n := range notes {
		noteResponses[i] = *n.ToNoteResponseDTO()
	}

	return noteResponses, nil
}

func (s *Service) GetNoteByID(id string) (*NoteResponseDTO, error) {
	note, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

	if note == nil {
		return nil, nil
	}

	return note.ToNoteResponseDTO(), nil
}
