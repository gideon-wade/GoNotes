package auth

type InMemUserRepository struct {
	users map[string]User
}

func NewInMemUserRepository() *InMemUserRepository {
	return &InMemUserRepository{
		users: make(map[string]User),
	}
}

func (r *InMemUserRepository) Save(user User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemUserRepository) GetByEmail(email string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, nil
}

func (r *InMemUserRepository) GetByID(id string) (*User, error) {
	if user, exists := r.users[id]; exists {
		return &user, nil
	}
	return nil, nil
}
