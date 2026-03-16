package note

type InMemNoteRepository struct {
}

func NewInMemNoteRepository() *InMemNoteRepository {
	return &InMemNoteRepository{}
}

var notes []Note = []Note{}

func (r *InMemNoteRepository) Save(note Note) error {
	notes = append(notes, note)
	return nil
}

func (r *InMemNoteRepository) GetAll() ([]Note, error) {
	return notes, nil
}

func (r *InMemNoteRepository) GetByID(id string) (*Note, error) {
	for _, n := range notes {
        if n.ID == id {
            return &n, nil
        }
    }

	return nil, nil	
}