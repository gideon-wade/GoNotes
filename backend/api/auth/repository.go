package auth

type UserRepository interface {
	Save(user User) error
	GetByEmail(email string) (*User, error)
	GetByID(id string) (*User, error)
}

type RefreshTokenRepository interface {
	Save(token RefreshToken) error
	GetByToken(token string) (*RefreshToken, error)
	DeleteByToken(token string) error
	DeleteAllTokensByUserID(userID string) error
}
