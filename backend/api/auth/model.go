package auth

import (
	"time"

	"github.com/google/uuid"
)

func NewRefreshToken(userID string, token string, expiresAt time.Time) *RefreshToken {
	return &RefreshToken{
		ID:        uuid.New().String(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}

type RefreshToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
