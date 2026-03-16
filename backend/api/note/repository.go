package note

type Repository interface {
	Save(note Note) error
	GetAll() ([]Note, error)
	GetByID(id string) (*Note, error)
}
