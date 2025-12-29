package note

type InMemNoteRepository struct {
}

func NewInMemNoteRepository() *InMemNoteRepository {
	return &InMemNoteRepository{}
}

var notes []Note = []Note{}

func (r *InMemNoteRepository) Save(note *Note) error {
	notes = append(notes, *note)
	return nil
}
