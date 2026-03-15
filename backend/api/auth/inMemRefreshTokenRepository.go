package auth

type InMemRefreshTokenRepository struct {
	tokens map[string]RefreshToken
}

func NewInMemRefreshTokenRepository() *InMemRefreshTokenRepository {
	return &InMemRefreshTokenRepository{
		tokens: make(map[string]RefreshToken),
	}
}

func (r *InMemRefreshTokenRepository) Save(token RefreshToken) error {
	r.tokens[token.ID] = token
	return nil
}

func (r *InMemRefreshTokenRepository) GetByToken(token string) (*RefreshToken, error) {
	for _, t := range r.tokens {
		if t.Token == token {
			return &t, nil
		}
	}
	return nil, nil
}

func (r *InMemRefreshTokenRepository) DeleteByToken(token string) error {
	for id, t := range r.tokens {
		if t.Token == token {
			delete(r.tokens, id)
			return nil
		}
	}
	return nil
}

func (r *InMemRefreshTokenRepository) DeleteAllByUserID(userID string) error {
	for id, t := range r.tokens {
		if t.UserID == userID {
			delete(r.tokens, id)
		}
	}
	return nil
}
