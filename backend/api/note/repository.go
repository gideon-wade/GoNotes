package note

type Repository interface {
	Save(note *Note) error
}
